package Impl

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:"port" envconfig:"SERVER_PORT"`
		Schema   string `json:"schema"`
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"database"`
	Redis struct {
		Address string `json:"address"`
		Db      int    `json:"db"`
	} `json:"redis"`
	DistMutex struct {
		CronTime       string `json:"cronTime"`
		CronMutex      string `json:"cronMutex"`
		ExpiryDuration int    `json:"expiryDuration"`
	} `json:"distMutex"`
	Host string `json:"host"`
	Port string `json:"port"`
}

func (c *Config) LoadConfiguration() {
	fileName := "config.json"
	configFile, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		return
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(c)
	err = configFile.Close()
	if err != nil {
		log.Println(err)
	}
}
