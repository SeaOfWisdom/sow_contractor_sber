package srv

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/SeaOfWisdom/sow_contractor/src/config"
	"github.com/SeaOfWisdom/sow_contractor/src/storage"
	lib "github.com/SeaOfWisdom/sow_proto/lib-srv"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// BscConnector ...
type Service struct {
	cfg    *config.Config
	logger *logrus.Entry
	client *ethclient.Client

	strSrv *storage.StorageService

	libSrv lib.LibraryServiceClient

	chainID      int64
	currentBlock uint64

	operatorAddress common.Address
	sowToken        common.Address
	sowLibrary      common.Address
}

// NewService ...
func NewService(logger *logrus.Logger, cfg *config.Config, strSrv *storage.StorageService, libSrv lib.LibraryServiceClient) *Service {
	client, err := ethclient.Dial(cfg.Provider)
	if err != nil {
		panic("new eth Client error")
	}

	privKey, err := getPrivateKey(cfg)
	if err != nil {
		panic(fmt.Sprintf("generate private key error, err: %s", err.Error()))
	}
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		panic(fmt.Sprintf("get public key error, err: %s", err.Error()))
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	if !bytes.Equal(common.HexToAddress(cfg.OperatorAddress).Bytes(), fromAddress.Bytes()) {
		panic(fmt.Sprintf(
			"relayer address supplied in config (%s) does not match mnemonic (%s)",
			cfg.OperatorAddress, fromAddress,
		))
	}
	// init token addresses
	return &Service{
		cfg:             cfg,
		chainID:         cfg.ChainID,
		logger:          logger.WithField("service", "SOW_CONTRACTOR"),
		client:          client,
		strSrv:          strSrv,
		libSrv:          libSrv,
		operatorAddress: common.HexToAddress(cfg.OperatorAddress),
		sowToken:        common.HexToAddress(cfg.SowTokenAddress),
		sowLibrary:      common.HexToAddress(cfg.SowLibraryAddress),
	}
}

func (s *Service) Start() {
	// get the start height
	currentHeight, err := s.strSrv.GetLatestKnownBlock()
	if err != nil {
		panic(fmt.Errorf("while get the latest block, err: %v", err))
	}
	s.currentBlock = currentHeight

	s.logger.Infoln("current block: ", s.currentBlock)

	go s.listenEvents()
}

func (s *Service) listenEvents() {
	latestBlock, err := s.GetCurrentBlock()
	if err != nil {
		panic(err)
	}
	toBlock := getToBlock(s.currentBlock, latestBlock)

	for {
		s.logger.Infof("fetch blocks: %d - %d", s.currentBlock, toBlock)
		if err := s.FetchLogsRange(big.NewInt(0).SetUint64(s.currentBlock), big.NewInt(0).SetUint64(toBlock)); err != nil {
			s.logger.Errorf("while fetch logs by range, err: %v", err)
			time.Sleep(15 * s.cfg.CheckerSleepTime)
			continue
		}
		// update the current height
		s.currentBlock = toBlock + 1
		if err := s.strSrv.UpdateLatestKnownBlock(toBlock); err != nil {
			s.logger.Errorf("while update the latest block in redis, err: %v", err)
			time.Sleep(15 * s.cfg.CheckerSleepTime)
			continue
		}

		latestBlock, err := s.GetCurrentBlock()
		if err != nil {
			time.Sleep(15 * s.cfg.CheckerSleepTime)
			continue
		}
		toBlock = getToBlock(s.currentBlock, latestBlock)

		if toBlock == latestBlock {
			time.Sleep(2 * s.cfg.CheckerSleepTime)
		}
		time.Sleep(s.cfg.CheckerSleepTime)
	}
}

func (s *Service) GetCurrentBlock() (uint64, error) {
	currentBlock, err := s.client.BlockNumber(context.Background())
	if err != nil {
		return 0, err
	}

	return currentBlock, nil
}

func (s *Service) CheckSentTxStatus(txHash string) {
	var count int64
	var err error
	for {
		var status bool
		if count > 5 {
			s.logger.Errorf("tx(%s) was NOT sent, err: %v", txHash, err)
			return
		}
		time.Sleep(5 * s.cfg.CheckerSleepTime)
		status, err = s.GetSentTxStatus(txHash)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				time.Sleep(2 * s.cfg.CheckerSleepTime)
				count++
				continue
			}
			s.logger.Errorf("tx(%s) was NOT sent, err: %v", txHash, err)
			return
		}

		if status {
			s.logger.Infof("Tx(%s) was sent successfully\n", txHash)
			return
		}
		count++
	}
}

// GetSentTxStatus ...
func (s *Service) GetSentTxStatus(hash string) (bool, error) {
	_, isPending, err := s.client.TransactionByHash(context.Background(), common.HexToHash(hash))
	if err != nil {
		return false, fmt.Errorf("NOT_FOUND")
	}
	if isPending {
		return false, nil
	}

	txReceipt, err := s.client.TransactionReceipt(context.Background(), common.HexToHash(hash))
	if err != nil {
		return false, err
	}

	if txReceipt.Status == types.ReceiptStatusFailed {
		return false, fmt.Errorf("SENT_FAILED")
	}

	return true, nil
}

// GetTransactor ...
func (s *Service) getTransactor() (auth *bind.TransactOpts, err error) {
	privateKey, err := getPrivateKey(s.cfg)
	if err != nil {
		return nil, err
	}

	nonce, err := s.client.PendingNonceAt(context.Background(), s.operatorAddress)
	if err != nil {
		return nil, err
	}

	auth, err = bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(s.chainID))
	if err != nil {
		return nil, err
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei

	return auth, nil
}

func getPrivateKey(config *config.Config) (*ecdsa.PrivateKey, error) {
	privKey, err := crypto.HexToECDSA(config.OperatorPrivateKey)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}

func getToBlock(currentBlock, latestBlock uint64) uint64 {
	if currentBlock+1000 < latestBlock {
		return currentBlock + 1000
	}
	return latestBlock
}
