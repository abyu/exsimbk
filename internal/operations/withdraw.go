package operations

import (
	"errors"
	"github.com/rwwae/simplebank/internal/accounts"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type Withdraw struct {
	amount float64
}

func NewWithdraw(amount float64) accounts.BalanceOperation {
	return &Withdraw{amount}
}

func (o *Withdraw) Perform(accountBalance float64) (float64, error) {
	if accountBalance < o.amount {
		return 0, ErrInsufficientFunds
	}
	return accountBalance - o.amount, nil
}