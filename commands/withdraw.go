package commands

import (
	"fmt"
	registry2 "oldbanksys/domain/registry"
)

func WithdrawAccount(accountId string, withdrawAmount float64) error {
	registry, err := registry2.GetRegistry()
	if err != nil {
		return err
	}

	acc, _ := registry.FindAccountById(accountId)
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
