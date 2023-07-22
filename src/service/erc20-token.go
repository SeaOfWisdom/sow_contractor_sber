package srv

import (
	"fmt"
	"math/big"

	erc20 "github.com/SeaOfWisdom/sow_contractor/src/service/abi/token"

	"github.com/ethereum/go-ethereum/common"
)

func (s *Service) burn(accountStr string, amount *big.Int) (hash string, err error) {
	if !common.IsHexAddress(accountStr) {
		err = fmt.Errorf("address %s is not HEX address", accountStr)

		return
	}

	account := common.HexToAddress(accountStr)

	auth, err := s.getTransactor()
	if err != nil {
		return
	}

	auth.GasLimit = defaultGasLimit

	instance, err := erc20.NewBond(s.sowToken, s.client)
	if err != nil {
		err = fmt.Errorf("error create erc20 bond instance: %v", err)

		return
	}

	tx, err := instance.Burn(auth, account, amount)
	if err != nil {
		err = fmt.Errorf("burn error: %v", err)

		return
	}

	hash = tx.Hash().Hex()

	return
}

func (s *Service) mint(accountStr string, amount *big.Int) (hash string, err error) {
	if !common.IsHexAddress(accountStr) {
		err = fmt.Errorf("address %s is not HEX address", accountStr)

		return
	}

	account := common.HexToAddress(accountStr)

	auth, err := s.getTransactor()
	if err != nil {
		return
	}

	//	auth.GasLimit = 100000 // 100 000

	instance, err := erc20.NewBond(s.sowToken, s.client)
	if err != nil {
		err = fmt.Errorf("error create erc20 bond instance: %v", err)

		return
	}

	tx, err := instance.Mint(auth, account, amount)
	if err != nil {
		err = fmt.Errorf("mint error: %v", err)

		return
	}

	hash = tx.Hash().Hex()

	return
}

func (s *Service) balanceOf(accountStr string) (balance *big.Int, err error) {
	if !common.IsHexAddress(accountStr) {
		err = fmt.Errorf("address %s is not HEX address", accountStr)

		return
	}

	instance, err := erc20.NewBond(s.sowToken, s.client)
	if err != nil {
		err = fmt.Errorf("error create erc20 bond instance: %v", err)

		return
	}

	return instance.BalanceOf(nil, common.HexToAddress(accountStr))
}
