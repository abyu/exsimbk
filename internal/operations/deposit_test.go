package operations

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"testing"
)

func TestDepositShouldAddTheGivenAmountToBalanceReturningTheNewBalance(t *testing.T) {
	deposit := NewDeposit(10)

	updatedBalance, _ := deposit.Perform(20)

	then.AssertThat(t, updatedBalance, is.EqualTo(float64(30)))
}
