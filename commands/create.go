package commands

import (
	"fmt"
	"oldbanksys/domain/registry"
)

func CreateAccount(name string, initialBalance float64) error {
	reg, err := registry.GetRegistry()
	if err != nil {
		return err
	}

	createdAccount := reg.CreateAccount(name, initialBalance)
	err = reg.WriteToLocalDbFile()
	if err != nil {
		return err
	}

	fmt.Println("Account created:")
	createdAccount.PrintToConsole()
	return err
}
