package commands

import (
	"fmt"
	registry2 "oldbanksys/domain/registry"
)

func ShowAccount(accountId string) error {
	registry, err := registry2.GetRegistry()
	if err != nil {
		return err
	}

	acc, _ := registry.FindAccountById(accountId)
	fmt.Println("=========================================================")
	acc.PrintToConsole()

	return nil
}
