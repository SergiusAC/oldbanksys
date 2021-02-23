package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oldbanksys/common"
	"oldbanksys/domain/registry"
)

func ShowAllAccounts() error {
	data, err := ioutil.ReadFile(common.DB_FILE_LOCATION)

	if err != nil {
		return err
	}

	reg := &registry.AccountRegistry{}
	err = json.Unmarshal(data, reg)

	if err != nil {
		return err
	}

	for _, account := range reg.Accounts {
		fmt.Println("=========================================================")
		account.PrintToConsole()
		fmt.Println()
	}

	return nil
}
