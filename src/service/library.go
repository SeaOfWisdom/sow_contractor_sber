package srv

import (
	"fmt"

	library "github.com/SeaOfWisdom/sow_contractor/src/service/abi/sow-library"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Service) AddParticipant(participantAddressStr string) (string, error) {
	if !common.IsHexAddress(participantAddressStr) {
		return "", fmt.Errorf("address %s is not HEX address", participantAddressStr)
	}
	account := common.HexToAddress(participantAddressStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}
	auth.GasLimit = 100000 // 100 000

	instance, err := library.NewLibrary(s.sowToken, s.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.AddParticipant(auth, account)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
