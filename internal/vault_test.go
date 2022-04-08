package internal

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/rwwae/simplebank/internal/accounts"
	"github.com/rwwae/simplebank/internal/operations"
	"testing"
)

func TestTotalBalanceIsTheSumOfAllAccountsBalances(t *testing.T) {
	vault := Vault{map[uint64]*accounts.Account{
		1: newAccount(1, 23),
		2: newAccount(2, 45),
		3: newAccount(3, 10),
	}}

	balance := vault.GetTotalBalance()

	then.AssertThat(t, balance, is.EqualTo(int64(78)))
}

func TestShouldReturnTheBalanceOfTheCustomerWithTheGivenAccountID(t *testing.T) {
	vault := Vault{map[uint64]*accounts.Account{
		1: newAccount(1, 23),
		2: newAccount(2, 45),
		3: newAccount(3, 10),
	}}

	balance, _ := vault.RetrieveBalance(2)

	then.AssertThat(t, balance, is.EqualTo(int64(45)))
}

func TestGetAccountBalanceShouldReturnInvalidAccountWhenNoAccountWithTheGivenAccountIDExists(t *testing.T) {
	vault := Vault{map[uint64]*accounts.Account{
		1: newAccount(1, 23),
		2: newAccount(2, 45),
		3: newAccount(3, 10),
	}}
	nonExistentAccountID := uint64(5)

	_, err := vault.RetrieveBalance(nonExistentAccountID)

	then.AssertThat(t, err, is.EqualTo(ErrInvalidAccount))
}

func TestDepositToAnAccountWithTheGivenAccountID(t *testing.T) {
	vault := Vault{map[uint64]*accounts.Account{
		1: newAccount(1, 0),
		2: newAccount(2, 10),
	}}
	firstAccountID := uint64(1)

	err := vault.Deposit(firstAccountID, 30)

	then.AssertThat(t, err, is.Nil())
	balance, _ := vault.RetrieveBalance(firstAccountID)
	then.AssertThat(t, balance, is.EqualTo(int64(30)))
}

func TestDepositingToAnAccountWithTheGivenAccountIDDoesNotAffectOtherAccounts(t *testing.T) {
	vault := Vault{map[uint64]*accounts.Account{
		1: newAccount(1, 0),
		2: newAccount(2, 10),
	}}

	err := vault.Deposit(1, 30)

	then.AssertThat(t, err, is.Nil())
	balance := vault.GetTotalBalance()
	then.AssertThat(t, balance, is.EqualTo(int64(40)))
}

func TestWithdrawingFromAnAccountWithGivenAccountIdReducesTheAccountsBalance(t *testing.T) {
	vault := Vault{map[uint64]*accounts.Account{
		1: newAccount(1, 40),
		2: newAccount(2, 10),
	}}
	firstAccountID := uint64(1)

	err := vault.Withdraw(firstAccountID, 20)

	then.AssertThat(t, err, is.Nil())
	balance, _ := vault.RetrieveBalance(firstAccountID)
	then.AssertThat(t, balance, is.EqualTo(int64(20)))
}

func TestWithdrawingReturnsInsufficientFundsErrWhenBalanceIsBelowWithdrawAmount(t *testing.T) {
	vault := Vault{map[uint64]*accounts.Account{
		1: newAccount(1, 20),
		2: newAccount(2, 10),
	}}
	firstAccountID := uint64(1)

	err := vault.Withdraw(firstAccountID, 40)

	then.AssertThat(t, err, is.EqualTo(operations.ErrInsufficientFunds))
	balance, _ := vault.RetrieveBalance(firstAccountID)
	then.AssertThat(t, balance, is.EqualTo(int64(20)))
}

func TestShouldReturnAListOfAllOperationsForTheGivenAccountID(t *testing.T) {
	account := accounts.NewAccount(1, 20, []accounts.BalanceOperation{
		operations.NewDeposit(10),
		operations.NewDeposit(30),
		operations.NewWithdraw(20),
	})
	vault := Vault{map[uint64]*accounts.Account{
		1: &account,
	}}
	firstAccountID := uint64(1)

	transactions, err := vault.GetAllTransactions(firstAccountID)

	then.AssertThat(t, err, is.Nil())
	then.AssertThat(t, len(transactions), is.EqualTo(3))
	then.AssertThat(t, transactions, is.EqualTo([]string{
		"Deposit 10", "Deposit 30", "Withdraw 20",
	}))
}

func newAccount(id uint64, balance int64) *accounts.Account {
	account := accounts.NewAccount(id, balance, []accounts.BalanceOperation{})
	return &account
}
