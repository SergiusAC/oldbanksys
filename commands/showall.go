package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oldbanksys/common"
	registry2 "oldbanksys/domain/registry"
)

func ShowAllAccounts() error {
	data, err := ioutil.ReadFile(common.DB_FILE_LOCATION)

	if err != nil {
		return err
	}

	registry := &registry2.AccountRegistry{}
	err = json.Unmarshal(data, registry)

	if err != nil {
		return err
	}

	for _, account := range registry.Accounts {
		fmt.Println("=========================================================")
		account.PrintToConsole()
		fmt.Println()
	}

	return nil
}
