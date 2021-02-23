package commands

import (
	"fmt"
	"oldbanksys/domain/registry"
)

func DeleteAccount(accountId string) error {
	reg, err := registry.GetRegistry()
	if err != nil {
		return err
	}

	deleted := reg.DeleteAccount(accountId)

	if deleted {
		err = reg.WriteToLocalDbFile()
		if err != nil {
			return err
		}
		fmt.Println("Account deleted")
	} else {
		fmt.Println("Account not found")
	}

	return nil
}
