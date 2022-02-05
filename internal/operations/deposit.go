package operations

import (
	"errors"
	"github.com/rwwae/simplebank/internal/accounts"
)

//ErrInvalidDepositAmount is returned when the amount is deposit is negative value
var ErrInvalidDepositAmount = errors.New("amount must be a positive value")

//Deposit operation with `amount` to deposit
type Deposit struct {
	amount float64
}

//NewDeposit ...
func NewDeposit(amount float64) accounts.BalanceOperation {
	return &Deposit{amount}
}

//Perform the deposit operation on the given accountBalance
func (o *Deposit) Perform(accountBalance float64) (float64, error) {
	if o.amount < 0 {
		return accountBalance, ErrInvalidDepositAmount
	}

	return o.amount + accountBalance, nil
}
