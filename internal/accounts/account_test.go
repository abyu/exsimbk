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

func TestShouldReturnTheListOfAllOperationsDoneSoFar(t *testing.T) {
	account := Account{id: 2, balanceInCents: 20}

	operation := &BalanceOperationTest{}
	operation2 := &BalanceOperationTest{}
	_, err := account.Apply(operation)
	then.AssertThat(t, err, is.Nil())

	_, err = account.Apply(operation2)
	then.AssertThat(t, err, is.Nil())

	balanceOperations := account.Transactions()
	then.AssertThat(t, len(balanceOperations), is.EqualTo(2))
	then.AssertThat(t, balanceOperations, is.EqualTo([]BalanceOperation{operation, operation2}))
}

type BalanceOperationTest struct {

}

func (b *BalanceOperationTest) Perform(accountBalance int64) (int64, error) {

	return 0, nil
}