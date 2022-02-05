package accounts

//BalanceOperation is contract for operations what work with account's balance
type BalanceOperation interface {
	Perform(accountBalance float64) (float64, error)
}

//BalanceOperationFunc is type alias that allows you use a single function as BalanceOperation without the need for
//creating a new struct.
type BalanceOperationFunc func(float64) (float64, error)

//Perform the operation on the `accountBalance`
func (boF BalanceOperationFunc) Perform(accountBalance float64) (float64, error) {
	return boF(accountBalance)
}
