package operations

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func TestDepositShouldAddTheGivenAmountToBalanceReturningTheNewBalance(t *testing.T) {
	deposit := NewDeposit(10)

	updatedBalance, _ := deposit.Perform(20)

	then.AssertThat(t, updatedBalance, is.EqualTo(int64(30)))
}

func TestDepositShouldReturnErrWhenTheGivenDepositAmountIsNegative(t *testing.T) {
	deposit := NewDeposit(-10)

	updatedBalance, err := deposit.Perform(20)

	then.AssertThat(t, err, is.EqualTo(ErrInvalidDepositAmount))
	then.AssertThat(t, updatedBalance, is.EqualTo(int64(20)))
}
