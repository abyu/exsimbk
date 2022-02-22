package accounts

//Account to hold the `balance` of the account with a unique id
type Account struct {
	id             uint64
	balanceInCents int64
}

//NewAccount ...
func NewAccount(id uint64, balance int64) Account {
	return Account{id, balance}
}

//Apply applies the `operation` and updates the balance from the resulting operation
func (a *Account) Apply(operation BalanceOperation) (int64, error) {

	updatedBalance, err := operation.Perform(a.balanceInCents)
	if err != nil {
		return 0, err
	}

	a.balanceInCents = updatedBalance
	return updatedBalance, nil
}

//Balance returns the current balance
func (a *Account) Balance() int64 {
	return a.balanceInCents
}

//ID returns the account id
func (a *Account) ID() uint64 {
	return a.id
}
