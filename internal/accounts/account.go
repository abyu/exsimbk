package accounts

//Account to hold the `balance` of the account with an unique id
type Account struct {
	id uint64
	balance float64
}

//BalanceOperation is contract for operations what work with account's balance
type BalanceOperation interface {
	Perform(accountBalance float64) (float64, error)
}

//BalanceOperationFunc is type alias that allows you use a single function as BalanceOperation without the need for
//creating a new struct.
type BalanceOperationFunc func(float64) (float64, error)

//Perform the operation on the `accountBalance`
func (boF BalanceOperationFunc) Perform(accountBalance float64)  (float64, error) {
	 return boF(accountBalance)
}

//NewAccount ...
func NewAccount(id uint64, balance float64) Account {
	return Account{id, balance}
}

//Apply applies the `operation` and updates the balance from the resulting operation
func (a *Account) Apply(operation BalanceOperation) (float64, error) {

	updatedBalance, err := operation.Perform(a.balance)
	if err != nil {
		return 0, err
	}

	a.balance = updatedBalance
	return updatedBalance, nil
}

//Balance returns the current balance
func (a *Account) Balance() float64 {
	return a.balance
}

//ID returns the account id
func (a *Account) ID() uint64 {
	return a.id
}
