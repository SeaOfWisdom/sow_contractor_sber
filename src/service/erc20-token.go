package srv

import (
	"fmt"
	"math/big"

	erc20 "github.com/SeaOfWisdom/sow_contractor/src/service/abi/token"
	"github.com/ethereum/go-ethereum/common"
)

func (s *Service) burn(accountStr string, amount *big.Int) (string, error) {
	if !common.IsHexAddress(accountStr) {
		return "", fmt.Errorf("address %s is not HEX address", accountStr)
	}
	account := common.HexToAddress(accountStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}
	auth.GasLimit = 100000 // 100 000

	instance, err := erc20.NewBond(s.sowToken, s.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.Burn(auth, account, amount)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) mint(accountStr string, amount *big.Int) (string, error) {
	if !common.IsHexAddress(accountStr) {
		return "", fmt.Errorf("address %s is not HEX address", accountStr)
	}
	account := common.HexToAddress(accountStr)

	auth, err := s.getTransactor()
	if err != nil {
		return "", err
	}
	//	auth.GasLimit = 100000 // 100 000

	instance, err := erc20.NewBond(s.sowToken, s.client)
	if err != nil {
		return "", err
	}

	tx, err := instance.Mint(auth, account, amount)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (s *Service) balanceOf(accountStr string) (*big.Int, error) {
	if !common.IsHexAddress(accountStr) {
		return nil, fmt.Errorf("address %s is not HEX address", accountStr)
	}

	instance, err := erc20.NewBond(s.sowToken, s.client)
	if err != nil {
		return nil, err
	}

	return instance.BalanceOf(nil, common.HexToAddress(accountStr))
}
