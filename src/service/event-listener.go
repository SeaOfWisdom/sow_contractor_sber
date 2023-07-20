package srv

import (
	"bytes"
	"context"
	"math/big"
	"strings"

	library "github.com/SeaOfWisdom/sow_contractor/src/service/abi/sow-library"
	lib "github.com/SeaOfWisdom/sow_proto/lib-srv"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/google/uuid"
)

var (
	purchasedPaperEvent = common.HexToHash("0x98fac159a44269757dea3c0bd3fec2bb4834543a7a2538d11bb167b67cb6cbd5")
)

func (s *Service) FetchLogsRange(fromHeight, toHeight *big.Int) (err error) {
	logs, err := s.client.FilterLogs(context.Background(), ethereum.FilterQuery{
		FromBlock: fromHeight,
		ToBlock:   toHeight,
		Addresses: []common.Address{s.sowLibrary},
	})
	if err != nil {
		return err
	}

	for _, log := range logs {
		var decodedEvent interface{}
		if bytes.Equal(log.Topics[0][:], purchasedPaperEvent[:]) {
			decodedEvent, err = s.decodePurchasedPaperEvent(log)
		}
		if err != nil {
			return err
		}

		if decodedEvent != nil {
			purchaseWorkEvent := decodedEvent.(*library.LibraryPaperPurchased)
			uuidStr, _ := uuid.FromBytes(purchaseWorkEvent.WorkId.Bytes())
			s.logger.Infof("new purchaseWork event: reader(%s) | workd: %s", purchaseWorkEvent.Reader, uuidStr.String())
			// send request to the sow-library
			if _, err := s.libSrv.MakeAsPurchased(context.Background(), &lib.MakeAsPurchasedRequest{
				ReaderAddress: purchaseWorkEvent.Reader.Hex(),
				WorkId:        uuidStr.String(),
			}); err != nil {
				if strings.Contains(err.Error(), "you have already purchased") {
					return nil
				}
				return err
			}
		}
	}
	return nil
}

func (s *Service) decodePurchasedPaperEvent(log types.Log) (*library.LibraryPaperPurchased, error) {
	var event library.LibraryPaperPurchased
	abi, _ := abi.JSON(strings.NewReader(library.LibraryABI))
	if err := abi.UnpackIntoInterface(&event, "PaperPurchased", log.Data); err != nil {
		s.logger.Errorf("while decoding StartWithdrawal, err: %v", err)
		return nil, err
	}
	event.Reader = common.BytesToAddress(log.Topics[1].Bytes())

	return &event, nil
}
