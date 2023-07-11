package srv

import (
	"fmt"
	"math/big"
)

func (s *Service) PurchaseWork(readerAddressStr, authorAddressStr, amountStr string) (readerTxHash string, authorTxHash string, err error) {
	amount, success := big.NewInt(0).SetString(amountStr, 10)
	if !success || amount.String() == "0" { // TODO > 0
		return "", "", fmt.Errorf("amount %s is wrong", amountStr)
	}

	// check the reader balance
	readerBalance, err := s.balanceOf(readerAddressStr)
	if err != nil {
		return "", "", err
	}

	if readerBalance.Cmp(amount) < 0 {
		return "", "", fmt.Errorf("insufficient reader balance %s", readerBalance.String())
	}

	readerTxHash, err = s.burn(readerAddressStr, amount)
	if err != nil {
		return "", "", err
	}

	rewards := big.NewInt(0).Quo(amount, big.NewInt(2))
	authorTxHash, err = s.mint(authorAddressStr, rewards)
	if err != nil {
		return "", "", err
	}
	return
}

func (s *Service) GetStatus() (string, error) {
	return "", nil
}
