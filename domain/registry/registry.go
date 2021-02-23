package registry

import (
	"encoding/json"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"oldbanksys/common"
	"oldbanksys/domain/account"
)

type AccountRegistry struct {
	Accounts []*account.Account
}

type IAccountRegistry interface {
	FindAccountById(accountId string) (*account.Account, int)
	CreateAccount(accountName string, initialBalance float64) *account.Account
	DeleteAccount(accountID string) bool
	WriteToLocalDbFile() error
}

func (registry *AccountRegistry) FindAccountById(id string) (*account.Account, int) {
	for idx, acc := range registry.Accounts {
		if acc.ID == id {
			return acc, idx
		}
	}
	return nil, 0
}

func (registry *AccountRegistry) CreateAccount(accountName string, initialBalance float64) *account.Account {
	newAcc := &account.Account{}
	genUuid := uuid.Must(uuid.NewV4())
	newAcc.ID = genUuid.String()
	newAcc.AccountName = accountName
	newAcc.Balance = initialBalance
	registry.Accounts = append(registry.Accounts, newAcc)
	return newAcc
}

func (registry *AccountRegistry) DeleteAccount(accountID string) bool {
	foundAccount, idx := registry.FindAccountById(accountID)
	if foundAccount != nil {
		registry.Accounts = append(registry.Accounts[:idx], registry.Accounts[idx+1:]...)
		return true
	}
	return false
}

func (registry *AccountRegistry) WriteToLocalDbFile() error {
	registryJson, err := json.Marshal(registry)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(common.DB_FILE_LOCATION, registryJson, 0644)

	return err
}
