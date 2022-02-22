package operations

import (
	"errors"
	"github.com/rwwae/simplebank/internal/accounts"
)

//ErrInvalidDepositAmount is returned when the amount to deposit is negative value
var ErrInvalidDepositAmount = errors.New("amount must be a positive value")

//Deposit operation with `amount` to deposit
type Deposit struct {
	amount int64
}

//NewDeposit ...
func NewDeposit(amount int64) accounts.BalanceOperation {
	return &Deposit{amount}
}

//Perform the deposit operation on the given accountBalance returning the updated balance
// or unchanged balance with relevant errors
func (o *Deposit) Perform(accountBalance int64) (int64, error) {
	if o.amount < 0 {
		return accountBalance, ErrInvalidDepositAmount
	}

	return o.amount + accountBalance, nil
}
