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
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

// BscConnector ...
type Service struct {
	cfg             *config.Config
	chainID         int64
	currentBlock    int64
	operatorAddress common.Address
	sowToken        common.Address
	logger          *logrus.Entry
	client          *ethclient.Client
}

// NewService ...
func NewService(logger *logrus.Logger, cfg *config.Config) *Service {
	client, err := ethclient.Dial(cfg.BscProvider)
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
	if !bytes.Equal(common.HexToAddress(cfg.BscOperatorAddress).Bytes(), fromAddress.Bytes()) {
		panic(fmt.Sprintf(
			"relayer address supplied in config (%s) does not match mnemonic (%s)",
			cfg.BscOperatorAddress, fromAddress,
		))
	}
	// init token addresses
	return &Service{
		cfg:             cfg,
		chainID:         cfg.BscChainID,
		logger:          logger.WithField("service", "SOW_CONTRACTOR"),
		client:          client,
		operatorAddress: common.HexToAddress(cfg.BscOperatorAddress),
		sowToken:        common.HexToAddress("0x800fc10bac18a0e2bd00fffaa760301cacec119b"), // TODO
	}
}

func (s *Service) GetCurrentBlock() *big.Int {
	return big.NewInt(int64(s.currentBlock))
}

// CheckSentTxStake checks sent tx status with GetSentTxStatus in the loop and updates stake
// in the storage
func (s *Service) CheckSentTxStatus(txHash string) {
	var count int64
	var err error
	for {
		var status bool
		if count > 5 {
			s.logger.Errorf("tx(%s) was NOT sent, err: %v", txHash, err)
			return
		}
		time.Sleep(5 * s.cfg.BscCheckerSleepTime)
		status, err = s.GetSentTxStatus(txHash)
		if err != nil {
			if strings.Contains(err.Error(), "not found") {
				time.Sleep(2 * s.cfg.BscCheckerSleepTime)
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
	privKey, err := crypto.HexToECDSA(config.BscOperatorPrivateKey)
	if err != nil {
		return nil, err
	}
	return privKey, nil
}
