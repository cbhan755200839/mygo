package configs

import (
	"encoding/json"
	"mygo/utils/logs"
	"mygo/utils/yamls"

	"go.uber.org/zap"
)

type YamlConfig struct {
	LogFile LogFile      `json:"log-file"`
	Server  ServerConfig `json:"server"`
}

type LogFile struct {
	AppPath string `json:"app-path"`
	ErrPath string `json:"err-path"`
}

type ServerConfig struct {
	Port int `json:"port"`
}

var Config = &YamlConfig{}

func InitConfig() (err error) {
	jsonStr, err := yamls.LoadYaml("config.yaml")
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(jsonStr), &Config)
	if err != nil {
		logs.Log.Fatal("json字符串转Config失败", zap.Error(err))
		return err
	}
	return nil
}
