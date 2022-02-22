package accounts

import (
	"errors"
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func TestShouldApplyTheGivenBalanceOperationToUpdateTheBalanceReturningTheNewBalance(t *testing.T) {
	account := Account{id: 1, balanceInCents: 10}

	newBalance, err := account.Apply(BalanceOperationFunc(func(balanceAmount int64) (int64, error) {
		return balanceAmount + 20, nil
	}))

	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, account.Balance(), is.EqualTo(int64(30)))
	then.AssertThat(t, newBalance, is.EqualTo(int64(30)))
}

func TestShouldNotApplyTheGivenBalanceOperationWhenItReturnsAnError(t *testing.T) {
	account := Account{id: 1, balanceInCents: 30}

	_, err := account.Apply(BalanceOperationFunc(func(balanceAmount int64) (int64, error) {
		return 0, errors.New("there is a problem")
	}))

	then.AssertThat(t, err, is.EqualTo(errors.New("there is a problem")))
	then.AssertThat(t, account.Balance(), is.EqualTo(int64(30)))
}
