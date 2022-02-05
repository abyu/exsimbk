Solution
----

Setup
==
1. Use golang 1.16, you can use either [gvm](https://github.com/moovweb/gvm#installing) or brew(for macs).
2. A goland IDE, any of the following should work,
   1. Jetbrains GoLand
   2. VsCode with go plugins
   3. Or else any text editor should be fine too(but will miss syntax highlighting and running tests from IDEs)

Design
==
1. Vault has a list of Account(s), is a logical collections on all Accounts in the bank.
   * Any operation that deals with Accounts in general, eg: Get an account, or a total balance.
   * There is an assumption that, the total balance sum of the balance of all accounts.
   * This exists to keep things simple for this solution, ideally we would have a ledger, an Account manager and so on. 
2. Account has an account id and Balance.
   * All operations that deal with a specific account in general, eg: Deposit, Withdraw.
   * A person/entity could have multiple accounts, but for now, we will associate an account to a person like Alice.
3. BalanceOperation is an operation that can be performed on the Account balance
   1. They abstract the logic how an operation such as deposit, withdrawal affects the balance along with a criteria.
   2. When applied to an account, the account simply updates it balance as returned by an operation.
   3. This allows us to the keep the responsibility of Account to a minimum and change/add new ways of deposit/withdrawal without changing the Account.


Notes:
1. Avoiding overflows, safe math,  use a Money class.

   