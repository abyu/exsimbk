package operations

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func TestWithdrawShouldReduceTheBalanceWithTheGivenAmountReturningTheNewBalance(t *testing.T) {
	withdraw := NewWithdraw(20)

	updatedBalance, err := withdraw.Perform(100)

	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, updatedBalance, is.EqualTo(float64(80)))
}

func TestWithdrawShouldReturnErrWhenTheGivenWithdrawAmountIsNegative(t *testing.T) {
	withdraw := NewWithdraw(-10)

	updatedBalance, err := withdraw.Perform(20)

	then.AssertThat(t, err, is.EqualTo(ErrInvalidWithdrawAmount))
	then.AssertThat(t, updatedBalance, is.EqualTo(float64(20)))
}

func TestWithdrawShouldReturnErrWhenTheWithdrawAmountIsMoreThanTheAccountBalance(t *testing.T) {
	withdraw := NewWithdraw(120)

	updatedBalance, err := withdraw.Perform(90)

	then.AssertThat(t, err, is.EqualTo(ErrInsufficientFunds))
	then.AssertThat(t, updatedBalance, is.EqualTo(float64(90)))
}
