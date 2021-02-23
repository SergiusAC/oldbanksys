package commands

import (
	"fmt"
	"oldbanksys/domain/registry"
)

func ShowAccount(accountId string) error {
	reg, err := registry.GetRegistry()
	if err != nil {
		return err
	}

	acc, _ := reg.FindAccountById(accountId)
	fmt.Println("=========================================================")
	acc.PrintToConsole()

	return nil
}
