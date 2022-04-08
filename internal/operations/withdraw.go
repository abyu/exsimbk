package operations

import (
	"errors"
	"fmt"
	"github.com/rwwae/simplebank/internal/accounts"
)

//ErrInsufficientFunds is returned when a withdrawal operation results in a negative balance
var ErrInsufficientFunds = errors.New("insufficient funds")

//ErrInvalidWithdrawAmount is returned when the withdrawal amount is a negative value
var ErrInvalidWithdrawAmount = errors.New("withdrawal amount must be a positive value")

//Withdraw operation with the `amount` to withdraw
type Withdraw struct {
	amount int64
}

//NewWithdraw ...
func NewWithdraw(amount int64) accounts.BalanceOperation {
	return &Withdraw{amount}
}

//Perform the withdrawal operation on the `accountBalance` returning the updated balance
// or unchanged balance with relevant errors.
func (o *Withdraw) Perform(accountBalance int64) (int64, error) {
	if o.amount < 0 {
		return accountBalance, ErrInvalidWithdrawAmount
	}

	if accountBalance-o.amount < 0 {
		return accountBalance, ErrInsufficientFunds
	}

	return accountBalance - o.amount, nil
}

func (o *Withdraw) String() string {
	return fmt.Sprintf("Withdraw %d", o.amount)
}