package commands

import (
	"fmt"
	"strconv"
)

func Dispatch(command string, args []string) error {
	switch command {
	case "showall":
		return dispatchShowAllCommand()

	case "create":
		return dispatchCreateCommand(args)

	case "delete":
		return dispatchDeleteCommand(args)

	case "deposit":
		return dispatchDepositCommand(args)

	case "withdraw":
		return dispatchWithdrawCommand(args)

	case "show":
		return dispatchShowCommand(args)

	default:
		return dispatchHelpCommand()
	}
}

func dispatchShowAllCommand() error {
	return ShowAllAccounts()
}

func dispatchCreateCommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("account name and initial balance not specified")
	}

	name := args[0]
	balance, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return err
	}

	return CreateAccount(name, balance)
}

func dispatchDeleteCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("account ID not provided")
	}

	accountId := args[0]
	return DeleteAccount(accountId)
}

func dispatchDepositCommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("account ID and deposit amount not specified")
	}

	accountId := args[0]
	depositAmount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return err
	}

	return DepositAccount(accountId, depositAmount)
}

func dispatchWithdrawCommand(args []string) error {
	if len(args) < 2 {
		return fmt.Errorf("account ID and withdraw amount not specified")
	}

	accountId := args[0]
	withdrawAmount, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return err
	}

	return WithdrawAccount(accountId, withdrawAmount)
}

func dispatchShowCommand(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("account ID not specified")
	}

	return ShowAccount(args[0])
}

func dispatchHelpCommand() error {
	ShowHelp()
	return nil
}
