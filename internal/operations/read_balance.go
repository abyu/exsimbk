package operations

import "github.com/rwwae/simplebank/internal/accounts"

var RetrieveBalance = accounts.BalanceOperationFunc(
	func(accountBalance float64) (float64, error) {
		return accountBalance, nil
	})
