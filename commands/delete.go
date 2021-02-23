package commands

import (
	"fmt"
	registry2 "oldbanksys/domain/registry"
)

func DeleteAccount(accountId string) error {
	registry, err := registry2.GetRegistry()
	if err != nil {
		return err
	}

	deleted := registry.DeleteAccount(accountId)

	if deleted {
		err = registry.WriteToLocalDbFile()
		if err != nil {
			return err
		}
		fmt.Println("Account deleted")
	} else {
		fmt.Println("Account not found")
	}

	return nil
}
