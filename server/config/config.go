package config

import (
	"encoding/json"
	"io/ioutil"
	"strings"
)

const mappingFile = "/etc/totoro/user_mapping.json"
const tokenFile = "/etc/totoro/token"

type UserMappingConfig struct {
	Mapping map[string]string `json:"mapping"`
}

func GetUserMapping() (*UserMappingConfig, error) {
	bytes, err := ioutil.ReadFile(mappingFile)
	if err != nil {
		return nil, err
	}

	result := &UserMappingConfig{}
	err = json.Unmarshal(bytes, result)
	if err != nil {
		return nil, err
	}
	return result, err
}

func GetGithubToken() (string, error) {
	bytes, err := ioutil.ReadFile(tokenFile)
	if err != nil {
		return "", err
	}

	return strings.TrimRight(string(bytes), "\n"), err
}
