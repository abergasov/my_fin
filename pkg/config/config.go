package config

import (
	"my_fin/backend/pkg/logger"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	DBHost    string `yaml:"db_host"`
	DBName    string `yaml:"db_name"`
	DBUser    string `yaml:"db_user"`
	DBPass    string `yaml:"db_pass"`
	DBPort    string `yaml:"db_port"`
	JWTKey    string `yaml:"jwt_key"`
	JWTLive   int64  `yaml:"jwt_live"`
	ProdEnv   bool   `yaml:"prod_env"`
	AppDomain string `yaml:"app_domain"`
	SSLEnable bool   `yaml:"ssl_enable"`
	IPApiKey  string `yaml:"ip_key"`
}

func InitConf() *AppConfig {
	path, err := os.Getwd()
	if err != nil {
		logger.Fatal("Can't locate current dir", err)
	}

	logger.Info("Current app dir: ", path)
	confFile := path + "/configs/app_conf.yml"
	confFile = filepath.Clean(confFile)
	logger.Info("Try open config file: ", confFile)

	file, errP := os.Open(confFile)
	if errP != nil {
		logger.Fatal("Can't open config file: "+confFile, errP)
	}
	defer file.Close()
	var cfg AppConfig
	decoder := yaml.NewDecoder(file)
	errD := decoder.Decode(&cfg)
	if errD != nil {
		logger.FatalDetailed("Invalid config file", errD, map[string]interface{}{"file_path": file})
	}

	return &cfg
}
