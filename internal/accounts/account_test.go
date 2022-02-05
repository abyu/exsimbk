package accounts

import (
	"errors"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func TestShouldApplyTheGivenBalanceOperationToUpdateTheBalanceReturningTheNewBalance(t *testing.T) {
	account := Account{id: 1, balance: 10}

	newBalance, err := account.Apply(BalanceOperationFunc(func(balanceAmount float64) (float64, error) {
		return balanceAmount + 20, nil
	}))

	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, account.Balance(), is.EqualTo(float64(30)))
	then.AssertThat(t, newBalance, is.EqualTo(float64(30)))
}

func TestShouldNotApplyTheGivenBalanceOperationWhenItReturnsAnError(t *testing.T) {
	account := Account{id: 1, balance: 30}

	_, err := account.Apply(BalanceOperationFunc(func(balanceAmount float64) (float64, error) {
		return 0, errors.New("there is a problem")
	}))

	then.AssertThat(t, err, is.EqualTo(errors.New("there is a problem")))
	then.AssertThat(t, account.Balance(), is.EqualTo(float64(30)))
}
