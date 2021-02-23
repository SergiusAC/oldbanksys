package commands

import (
	"fmt"
	"oldbanksys/domain/registry"
)

func DepositAccount(accountId string, depositAmount float64) error {
	reg, err := registry.GetRegistry()
	if err != nil {
		return err
	}

	account, _ := reg.FindAccountById(accountId)
	if account == nil {
		return fmt.Errorf("account with ID %s not found", accountId)
	}

	err = account.Deposit(depositAmount)
	if err != nil {
		return err
	}

	err = reg.WriteToLocalDbFile()
	if err != nil {
		return err
	}

	fmt.Println("Account deposited:")
	account.PrintToConsole()

	return nil
}
