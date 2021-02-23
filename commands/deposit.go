package commands

import (
	"fmt"
	registry2 "oldbanksys/domain/registry"
)

func DepositAccount(accountId string, depositAmount float64) error {
	registry, err := registry2.GetRegistry()
	if err != nil {
		return err
	}

	account, _ := registry.FindAccountById(accountId)
	if account == nil {
		return fmt.Errorf("account with ID %s not found", accountId)
	}

	err = account.Deposit(depositAmount)
	if err != nil {
		return err
	}

	err = registry.WriteToLocalDbFile()
	if err != nil {
		return err
	}

	fmt.Println("Account deposited:")
	account.PrintToConsole()

	return nil
}
