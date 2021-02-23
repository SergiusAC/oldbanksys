package commands

import (
	"fmt"
	registry2 "oldbanksys/domain/registry"
)

func CreateAccount(name string, initialBalance float64) error {
	registry, err := registry2.GetRegistry()
	if err != nil {
		return err
	}

	createdAccount := registry.CreateAccount(name, initialBalance)
	err = registry.WriteToLocalDbFile()
	if err != nil {
		return err
	}

	fmt.Println("Account created:")
	createdAccount.PrintToConsole()
	return err
}
