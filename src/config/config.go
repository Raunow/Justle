package config

import (
	"encoding/json"
	"io/ioutil"
)

type Secrets struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Url      string `json:"url"`
}

func ReadConfig(configPaths ...string) (secrets Secrets) {
	configPath := ""
	if len(configPaths) == 0 {
		configPath = "./secrets.json"
	}

	bytes, err := ioutil.ReadFile(configPath)

	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(bytes, &secrets)
	if err != nil {
		panic(err)
	}

	return secrets
}
