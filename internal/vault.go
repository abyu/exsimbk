package internal

import (
	"errors"
	"github.com/rwwae/simplebank/internal/accounts"
	"github.com/rwwae/simplebank/internal/operations"
)

type Vault struct {
	accounts map[int]*accounts.Account
}

var ErrInvalidAccount = errors.New("invalid account")

func (v *Vault) GetTotalBalance() float64 {
	totalBalance := float64(0)
	for _, account:= range v.accounts {
		totalBalance += account.Balance()
	}

	return totalBalance
}

func (v *Vault) Deposit(accountID int, amount float64) error {
	depositOperation := operations.NewDeposit(amount)

	_, err := v.performOnValidAccount(accountID, depositOperation)

	return err
}

func (v *Vault) Withdraw(accountID int, amount float64) error {
	withdrawOperation := operations.NewWithdraw(amount)

	_, err := v.performOnValidAccount(accountID, withdrawOperation)

	return err
}

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

