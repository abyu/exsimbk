package operations

import (
	"github.com/rwwae/simplebank/internal/accounts"
)

type Deposit struct {
	amount float64
}

func NewDeposit(amount float64) accounts.BalanceOperation {
	return &Deposit{amount}
}

func (o *Deposit) Perform(accountBalance float64) (float64, error) {
	return o.amount + accountBalance, nil
}
