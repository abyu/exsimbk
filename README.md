Solution
----

Setup
==
1. Use golang 1.16, you can use either [gvm](https://github.com/moovweb/gvm#installing) or brew(for macs) to install it.
2. A goland IDE, any of the following should work,
   1. Jetbrains GoLand
   2. VsCode with go plugins
   3. Or else any text editor should be fine too(but will miss syntax highlighting and running tests from IDEs)

Running tests
==
1. Run `make test` to run all tests.
   1. Alternatively, run `go get -d -v ./...` to install dependencies.
   2. And run `go test -v ./...` from the project root folder to run the test, if make is not available.
2. The tests can also be run from the IDE, both goland and vscode support it.

Dir structure
==
1. All files are under `internal` folder.
2. The tests files are alongside the solution code.
3. `scenario_test.go` contains high level tests, start here, this also contains the scenario mentioned in the exercise statement along with couple others.

Design
==
1. Vault has a list of Account(s), is a logical collections of all Accounts in the bank.
   * The accounts are mapped to their ID for easy retrieval, ideally a persistence layer would be useful here.
   * This provides the API to access Account information and perform actions on individual account.
   * There is an assumption that, the total balance is sum of the balances of all accounts present in the vault.
   * This exists to keep things simple for this solution, ideally we would have a ledger, an Account manager and so on. 
2. Account has an account id and Balance.
   * All operations that deal with a specific account, eg: Deposit, Withdraw are performed here.
   * For changes to Balances, any operation defined as BalanceOperation can be applied to the account.
   * A person/entity could have multiple accounts, but for now, we will associate an account to a person like Alice.
3. BalanceOperation is an operation that can be performed on the Account balance
   1. This is a contract that defines how various operations such as Deposit and withdrawal can be defined and how they interact with the account.
   2. They allow abstraction of the logic on how an operation such as deposit, withdrawal affects the balance along with any criteria.
   3. When applied to an account, the account simply updates its balance to the value returned by an operation.
   4. This allows us to the keep the responsibility of Account to a minimum and change/add new ways of deposit/withdrawal without modifying the Account implementation.

Few trade-offs:
==
1. The balance operations does not avoid any overflow errors, I have used float64 for the amounts, that has sufficiently higher range.
2. Vault class creates new instance of Deposit and Withdraw operation on line 25 and 33, this might be a problem for testing.
   1. The vault cannot be tested as single unit right now because of this.
   2. This can be fixed by using a factory and passing it as a dependency to Vault, this would allow us to mock the operations.
3. I haven't used any Behavioral test frameworks, just added some helpers methods to make the scenario tests easily readable.

   