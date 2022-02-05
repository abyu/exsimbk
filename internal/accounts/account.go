package accounts

//Account to hold the `balance` of the account with a unique id
type Account struct {
	id      uint64
	balance float64
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
