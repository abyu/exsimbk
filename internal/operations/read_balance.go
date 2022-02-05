package operations

import "github.com/rwwae/simplebank/internal/accounts"

//RetrieveBalance is pass through operation that returns the accountBalance unchanged.
var RetrieveBalance = accounts.BalanceOperationFunc(
	func(accountBalance float64) (float64, error) {
		return accountBalance, nil
	})
