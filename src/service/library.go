package srv

import (
	"context"
	"fmt"
	"math/big"

	library "github.com/SeaOfWisdom/sow_contractor/src/service/abi/sow-library"

	"github.com/ethereum/go-ethereum/common"
)

const defaultGasLimit = 100_000

func (s *Service) AddParticipant(ctx context.Context, participantAddressStr string) (hash string, err error) {
	s.logger.Debugf("AddParticipant: participant address - %s", participantAddressStr)

	if !common.IsHexAddress(participantAddressStr) {
		err = fmt.Errorf("AddParticipant: address %s is not HEX address", participantAddressStr)

		return
	}

	account := common.HexToAddress(participantAddressStr)

	auth, err := s.getTransactor()
	if err != nil {
		return
	}

	auth.GasLimit = defaultGasLimit

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		err = fmt.Errorf("AddParticipant: %v", err)

		return
	}

	tx, err := instance.AddParticipant(auth, account)
	if err != nil {
		err = fmt.Errorf("AddParticipant: %v", err)

		return
	}

	hash = tx.Hash().Hex()

	return
}

func (s *Service) PublishPaper(
	authorAddresses []string,
	name,
	paperURI,
	paperIdStr, paperPriceStr string,
) (string, error) {
	s.logger.Debugf("PublishPaper: authorAddresses - %s\nname - %s, paperURI - %s\npaperIdStr - %s, paperPriceStr - %s",
		authorAddresses, name, paperURI, paperIdStr, paperPriceStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return "", err
	}

	var authors []common.Address
	for _, author := range authorAddresses {
		authors = append(authors, common.HexToAddress(author))
	}
	paperId, _ := big.NewInt(0).SetString(paperIdStr, 10)
	paperPrice, _ := big.NewInt(0).SetString(paperPriceStr, 10)

	tx, err := instance.PublishPaper(
		auth,
		authors,
		name,
		paperURI,
		paperId,
		paperPrice,
	)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) MakeAuthor(participantAddressStr string) (string, error) {
	s.logger.Debugf("MakeAuthor: participant address - %s", participantAddressStr)

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
	s.logger.Debugf("MakeReviewer: participant address - %s", participantAddressStr)

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
	s.logger.Debugf("MakeAdmin: participant address - %s", participantAddressStr)

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

func (s *Service) GetPaperById(paperIdStr string) (common.Address, error) {
	s.logger.Debugf("GetPaperById: paperId - %s", paperIdStr)

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return common.Address{}, err
	}
	paperId, _ := big.NewInt(0).SetString(paperIdStr, 10)

	return instance.GetPaperById(nil, paperId)
}

func (s *Service) GetParticipantRole(participantAddressStr string) (uint8, error) {
	s.logger.Debugf("GetParticipantRole: participantAddress - %s", participantAddressStr)

	if !common.IsHexAddress(participantAddressStr) {
		return 0, fmt.Errorf("address %s is not HEX address", participantAddressStr)
	}
	account := common.HexToAddress(participantAddressStr)

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return 0, err
	}

	return instance.GetRole(nil, account)
}

func (s *Service) IsAuthor(participantAddressStr string) (bool, error) {
	s.logger.Debugf("GetParticipantRole: participantAddress - %s", participantAddressStr)

	if !common.IsHexAddress(participantAddressStr) {
		return false, fmt.Errorf("address %s is not HEX address", participantAddressStr)
	}
	account := common.HexToAddress(participantAddressStr)

	instance, err := library.NewLibrary(s.sowLibrary, s.client)
	if err != nil {
		return false, err
	}

	return instance.IsAuthor(nil, account)
}
