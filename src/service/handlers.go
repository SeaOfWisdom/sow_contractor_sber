package srv

import (
	"fmt"
	"math/big"
)

func (s *Service) PurchaseWork(readerAddressStr, authorAddressStr, amountStr string) (readerTxHash string, authorTxHash string, err error) {
	s.logger.Debugf("PurchaseWork: reader address - %s, author address - %s, amount - %s", readerAddressStr, authorAddressStr, amountStr)

	amount, success := big.NewInt(0).SetString(amountStr, 10)
	if !success || amount.String() == "0" { // TODO > 0
		err = fmt.Errorf("PurchaseWork: amount %s is wrong", amountStr)

		return
	}

	// check the reader balance
	readerBalance, err := s.balanceOf(readerAddressStr)
	if err != nil {
		return
	}

	if readerBalance.Cmp(amount) < 0 {
		err = fmt.Errorf("PurchaseWork: insufficient reader balance %s", readerBalance.String())

		return
	}

	readerTxHash, err = s.burn(readerAddressStr, amount)
	if err != nil {
		return
	}

	rewards := big.NewInt(0).Quo(amount, big.NewInt(2))
	authorTxHash, err = s.mint(authorAddressStr, rewards)
	if err != nil {
		return
	}

	return
}

func (s *Service) Faucet(addressStr, amountStr string) (hash string, err error) {
	s.logger.Debugf("Faucet: address - %s, amount - %s", addressStr, amountStr)

	amount, success := big.NewInt(0).SetString(amountStr, 10)
	if !success || amount.String() == "0" { // TODO > 0
		err = fmt.Errorf("Faucet: amount %s is wrong", amountStr)

		return
	}

	return s.mint(addressStr, amount)
}

func (s *Service) GetStatus() (status string, err error) {
	return
}
