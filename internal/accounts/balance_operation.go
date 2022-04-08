package accounts

//BalanceOperation is contract for operations what work with account's balance
type BalanceOperation interface {
	Perform(accountBalance int64) (int64, error)
	String() string
}

//BalanceOperationFunc is type alias that allows you use a single function as BalanceOperation without the need for
//creating a new struct.
type BalanceOperationFunc func(int64) (int64, error)

//Perform the operation on the `accountBalance`
func (boF BalanceOperationFunc) Perform(accountBalance int64) (int64, error) {
	return boF(accountBalance)
}

func (boF BalanceOperationFunc) String() string {
	return ""
}
