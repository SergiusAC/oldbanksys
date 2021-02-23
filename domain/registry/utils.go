package registry

import (
	"encoding/json"
	"io/ioutil"
	"oldbanksys/common"
)

func InitEmptyRegistry() error {
	registry := &AccountRegistry{}
	registryJson, err := json.Marshal(registry)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(common.DB_FILE_LOCATION, registryJson, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetRegistryIfExists() (*AccountRegistry, error) {
	data, err := ioutil.ReadFile(common.DB_FILE_LOCATION)
	if err != nil {
		return nil, err
	}

	registry := &AccountRegistry{}

	err = json.Unmarshal(data, registry)
	if err != nil {
		return nil, err
	}

	return registry, nil
}

func GetRegistry() (*AccountRegistry, error) {
	registry, err := GetRegistryIfExists()

	if err != nil {
		err = InitEmptyRegistry()
		if err != nil {
			return nil, err
		}

		return GetRegistryIfExists()
	}

	return registry, nil
}
