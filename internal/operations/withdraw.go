package operations

import (
	"errors"
	"github.com/rwwae/simplebank/internal/accounts"
)

//ErrInsufficientFunds is returned when a withdrawal operation results in a negative balance
var ErrInsufficientFunds = errors.New("insufficient funds")

//ErrInvalidWithdrawAmount is returned when the withdrawal amount is a negative value
var ErrInvalidWithdrawAmount = errors.New("withdrawal amount must be a positive value")

//Withdraw operation with the `amount` to withdraw
type Withdraw struct {
	amount float64
}

//NewWithdraw ...
func NewWithdraw(amount float64) accounts.BalanceOperation {
	return &Withdraw{amount}
}

//Perform the withdrawal operation on the `accountBalance` returning the updated balance
// or unchanged balance with relevant errors.
func (o *Withdraw) Perform(accountBalance float64) (float64, error) {
	if o.amount < 0 {
		return accountBalance, ErrInvalidWithdrawAmount
	}

	if accountBalance-o.amount < 0 {
		return accountBalance, ErrInsufficientFunds
	}

	return accountBalance - o.amount, nil
}
