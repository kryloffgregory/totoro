package config

import (
	"encoding/json"
	"os"
)

const configFile = "/etc/totoro/user_mapping.json"

type UserMappingConfig struct {
	Mapping map[string]string `json:"Mapping"`
}

func GetConfig()  (*UserMappingConfig, error) {
	bytes, err:=os.ReadFile(configFile)
	if err!=nil {
		return nil, err
	}

	result :=&UserMappingConfig{}
	err = json.Unmarshal(bytes, result)
	if err!=nil {
		return nil, err
	}
	return result, err
}
