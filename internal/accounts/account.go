package accounts

type Account struct {
	id uint64
	balance float64
}

type BalanceOperation interface {
	Perform(accountBalance float64) (float64, error)
}

type BalanceOperationFunc func(float64) (float64, error)

func (boF BalanceOperationFunc) Perform(amount float64)  (float64, error) {
	 return boF(amount)
}

func NewAccount(id uint64, balance float64) Account {
	return Account{id, balance}
}

func (a *Account) Apply(operation BalanceOperation) (float64, error) {

	updatedBalance, err := operation.Perform(a.balance)
	if err != nil {
		return 0, err
	}

	a.balance = updatedBalance
	return updatedBalance, nil
}

func (a *Account) Balance() float64 {
	return a.balance
}
