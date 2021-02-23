package commands

import (
	"fmt"
	"oldbanksys/domain/registry"
)

func WithdrawAccount(accountId string, withdrawAmount float64) error {
	reg, err := registry.GetRegistry()
	if err != nil {
		return err
	}

	acc, _ := reg.FindAccountById(accountId)
	if acc == nil {
		return fmt.Errorf("account with ID %s not found", accountId)
	}

	err = acc.Withdraw(withdrawAmount)
	if err != nil {
		return err
	}

	fmt.Println("Account withdraw succeed")
	return nil
}
