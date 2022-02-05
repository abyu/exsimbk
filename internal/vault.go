package internal

import (
	"errors"
	"github.com/rwwae/simplebank/internal/accounts"
	"github.com/rwwae/simplebank/internal/operations"
)

//Vault contains a list of all accounts mapped ot their ids for easy retrieval
type Vault struct {
	accounts map[int]*accounts.Account
}

//ErrInvalidAccount is returned when accessing an account with a non-existent account id
var ErrInvalidAccount = errors.New("invalid account")

//GetTotalBalance returns the total all account balances
func (v *Vault) GetTotalBalance() float64 {
	totalBalance := float64(0)
	for _, account:= range v.accounts {
		totalBalance += account.Balance()
	}

	return totalBalance
}

//Deposit deposits the 'amount` to the account with 'accountID' when present returning an error otherwise.
func (v *Vault) Deposit(accountID int, amount float64) error {
	depositOperation := operations.NewDeposit(amount)

	_, err := v.performOnValidAccount(accountID, depositOperation)

	return err
}

//Withdraw withdraws the 'amount` from the account with 'accountID' when present with sufficient balance
// returning an error otherwise.
func (v *Vault) Withdraw(accountID int, amount float64) error {
	withdrawOperation := operations.NewWithdraw(amount)

	_, err := v.performOnValidAccount(accountID, withdrawOperation)

	return err
}

//RetrieveBalance returns the current balance of the account with `accountID` when present returning an error otherwise
func (v *Vault) RetrieveBalance(accountID int) (float64, error) {
	newBalance, err := v.performOnValidAccount(accountID, operations.RetrieveBalance)
	if err != nil {
		return 0, err
	}

	return newBalance, nil
}

func (v *Vault) performOnValidAccount(accountID int, operation accounts.BalanceOperation) (float64, error) {
	if account, ok := v.accounts[accountID]; ok {
		return account.Apply(operation)
	}

	return 0, ErrInvalidAccount
}

