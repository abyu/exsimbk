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