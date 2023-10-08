package configs

import (
	"Nik/pkg/models"
	"encoding/json"
	"os"
)

func InitConfigs() (config *models.Configs, err error) {
	bytes, err := os.ReadFile("./internal/configs/configs.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &config)
	if err != nil {
		return nil, err
	}
	return config, nil
}
