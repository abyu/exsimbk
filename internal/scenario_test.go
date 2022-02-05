package internal

import (
	"github.com/corbym/gocrest/is"
	"github.com/corbym/gocrest/then"
	"github.com/rwwae/simplebank/internal/accounts"
	"github.com/rwwae/simplebank/internal/operations"
	"testing"
)

// Test scenarios

//Example test scenario mentioned in exercise
func TestScenarioOne(t *testing.T) {
	alice := newAccount(1, 0)
	vault := Vault{accounts: map[uint64]*accounts.Account{1: alice}}

	//Alice deposits $30
	_ = vault.Deposit(alice.ID(), 30)
	accountBalanceMustBe(t, vault, alice.ID(), 30)
	bankBalanceMustBe(t, vault, 30)

	//Alice withdraws $20
	_ = vault.Withdraw(alice.ID(), 20)
	accountBalanceMustBe(t, vault, alice.ID(), 10)
	bankBalanceMustBe(t, vault, 10)

	//Alice fails withdraws another $11
	err := vault.Withdraw(alice.ID(), 11)
	then.AssertThat(t, err, is.EqualTo(operations.ErrInsufficientFunds))
	accountBalanceMustBe(t, vault, alice.ID(), 10)
	bankBalanceMustBe(t, vault, 10)
}

func TestScenarioTwo(t *testing.T) {
	alice := newAccount(1, 0)
	bob := newAccount(2, 0)
	vault := Vault{accounts: map[uint64]*accounts.Account{
		alice.ID(): alice,
		bob.ID():   bob,
	}}

	//Alice deposits $30, Bob deposits $50
	_ = vault.Deposit(alice.ID(), 30)
	_ = vault.Deposit(bob.ID(), 50)
	bankBalanceMustBe(t, vault, 80)

	//Alice fails to withdraw $80
	err := vault.Withdraw(alice.ID(), 80)
	then.AssertThat(t, err, is.EqualTo(operations.ErrInsufficientFunds))
	bankBalanceMustBe(t, vault, 80)
}

func TestScenarioThree(t *testing.T) {
	alice := newAccount(1, 0)
	bob := newAccount(2, 0)
	vault := Vault{accounts: map[uint64]*accounts.Account{
		alice.ID(): alice,
		bob.ID():   bob,
	}}

	//Alice deposits $30, Bob deposits $50
	_ = vault.Deposit(alice.ID(), 30)
	_ = vault.Deposit(bob.ID(), 50)
	bankBalanceMustBe(t, vault, 80)

	//Bob withdraws $40, Alice's balance is unchanged
	_ = vault.Withdraw(bob.ID(), 40)
	accountBalanceMustBe(t, vault, alice.ID(), 30)
	accountBalanceMustBe(t, vault, bob.ID(), 10)
	bankBalanceMustBe(t, vault, 40)
}

func accountBalanceMustBe(t *testing.T, vault Vault, id uint64, expected float64) {
	balanceForAlice, _ := vault.RetrieveBalance(id)
	then.AssertThat(t, balanceForAlice, is.EqualTo(expected))
}

func bankBalanceMustBe(t *testing.T, vault Vault, expected float64) {
	then.AssertThat(t, vault.GetTotalBalance(), is.EqualTo(expected))
}
