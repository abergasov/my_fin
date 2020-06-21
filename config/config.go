package config

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type AppConfig struct {
	DBHost  string
	DBName  string
	DBUser  string
	DBPass  string
	DBPort  string
	JWTKey  string
	JWTLive int64
	ProdEnv bool
}

var conf *AppConfig

func InitConf() *AppConfig {
	parsedConf := readConf()
	mode := getVariableOrDefault(parsedConf, "APP_ENV", "DEV")
	jwtLiveString := getVariableOrDefault(parsedConf, "JWT_LIVE_TIME", "")
	jwtLive, err := strconv.ParseInt(jwtLiveString, 10, 64)
	if err != nil {
		jwtLive = 0
	}
	return &AppConfig{
		DBHost:  getVariableOrDefault(parsedConf, "DB_HOST", ""),
		DBName:  getVariableOrDefault(parsedConf, "DB_NAME", ""),
		DBUser:  getVariableOrDefault(parsedConf, "DB_USER", ""),
		DBPass:  getVariableOrDefault(parsedConf, "DB_PASS", ""),
		DBPort:  getVariableOrDefault(parsedConf, "DB_PORT", ""),
		JWTKey:  getVariableOrDefault(parsedConf, "JWT_KEY", ""),
		JWTLive: jwtLive,
		ProdEnv: mode == "PROD",
	}
}

func readConf() *map[string]string {
	path, err := os.Getwd()
	if err != nil {
		log.Println("Can't locate current dir", err)
		panic(err.Error())
	}
	confFile := path + "/config/.env"
	file, errP := os.Open(confFile)
	if errP != nil {
		log.Print("Can't open config config file", confFile, errP)
		panic(errP.Error())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	tmpConf := map[string]string{}
	for {
		line, err := reader.ReadString('\n')

		equal := strings.Index(line, "=")
		if equal == -1 {
			continue
		}

		if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
			if len(line) <= equal {
				continue
			}
			tmpConf[key] = strings.TrimSpace(line[equal+1:])
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err.Error())
		}
	}
	return &tmpConf
}

func getVariableOrDefault(tmpConf *map[string]string, name string, defaultValue string) string {
	for key, val := range *tmpConf {
		if key == name {
			return val
		}
	}
	return defaultValue
}
