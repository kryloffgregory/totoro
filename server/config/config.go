package config

import (
	"encoding/json"
	"os"
)

const mappingFile = "/etc/totoro/user_mapping.json"
const tokenFile = "/etc/totoro/token"

type UserMappingConfig struct {
	Mapping map[string]string `json:"Mapping"`
}

func GetUserMapping()  (*UserMappingConfig, error) {
	bytes, err:=os.ReadFile(mappingFile)
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

func GetGithubToken() (string, error){
	bytes, err:=os.ReadFile(tokenFile)
	if err!=nil {
		return "", err
	}

	result :=""
	err = json.Unmarshal(bytes, &result)
	if err!=nil {
		return "", err
	}
	return result, err
}
