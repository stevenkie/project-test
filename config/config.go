package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/stevenkie/project-test/util"
)

type Config struct {
	DB     DBConfig     `json:"db"`
	Redis  RedisConfig  `json:"redis"`
	Server ServerConfig `json:"server"`
}

type DBConfig struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Name    string `json:"name"`
	SSLMode string `json:"sslmode"`
}

type RedisConfig struct {
	Host string `json:"host"`
}

type ServerConfig struct {
	Port   string `json:"port"`
	Secret string `json:"secret"`
}

func GetConfig() *Config {
	var byteValue []byte
	var conf *Config
	jsonFile, err := os.Open(fmt.Sprintf("./config/%s.json", util.GetEnv()))
	if err != nil {
		log.Fatalln(err)
	}
	defer jsonFile.Close()
	byteValue, err = ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal([]byte(byteValue), &conf)
	if err != nil {
		log.Fatalln(err)
	}
	return conf
}
