package commands

import "fmt"

func ShowHelp() {
	instruction := `OldBankSys commands
showall -- show information about all accounts
show <account_id> -- show information about account specified by <account_id>
create <account_name> <initial_balance> -- create new account with provided account name and initial balance
delete <account_id> -- delete an account by provided account ID
deposit <account_id> <amount> -- deposit an account by provided account ID and deposit amount
withdraw <account_id> <amount> -- withdraw an account by provided account ID and amount to withdraw
`
	fmt.Print(instruction)
}
