package srv

import (
	"context"
	"fmt"

	library "github.com/SeaOfWisdom/sow_contractor/src/service/abi/sow-library"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Service) AddParticipant(ctx context.Context, participantAddressStr string) (string, error) {
	if !common.IsHexAddress(participantAddressStr) {
		return "", fmt.Errorf("address %s is not HEX address", participantAddressStr)
	}
	account := common.HexToAddress(participantAddressStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}
	auth.GasLimit = 100000 // 100 000

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.AddParticipant(auth, account)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) PublishPaper(participantAddressStr string) (string, error) {
	if !common.IsHexAddress(participantAddressStr) {
		return "", fmt.Errorf("address %s is not HEX address", participantAddressStr)
	}
	account := common.HexToAddress(participantAddressStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}
	auth.GasLimit = 100000 // 100 000

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.AddParticipant(auth, account)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) MakeAuthor(participantAddressStr string) (string, error) {
	if !common.IsHexAddress(participantAddressStr) {
		return "", fmt.Errorf("address %s is not HEX address", participantAddressStr)
	}
	account := common.HexToAddress(participantAddressStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}
	auth.GasLimit = 100000 // 100 000

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.MakeAuthor(auth, account)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) MakeReviewer(participantAddressStr string) (string, error) {
	if !common.IsHexAddress(participantAddressStr) {
		return "", fmt.Errorf("address %s is not HEX address", participantAddressStr)
	}
	account := common.HexToAddress(participantAddressStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}
	auth.GasLimit = 100000 // 100 000

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.MakeReviewer(auth, account)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) MakeAdmin(participantAddressStr string) (string, error) {
	if !common.IsHexAddress(participantAddressStr) {
		return "", fmt.Errorf("address %s is not HEX address", participantAddressStr)
	}
	account := common.HexToAddress(participantAddressStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}
	auth.GasLimit = 100000 // 100 000

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.MakeAdmin(auth, account)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}
