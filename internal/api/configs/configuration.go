package configs

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server    Server    `yaml:"server"`
	Messaging Messaging `yaml:"messaging"`
	Database  Database  `yaml:"database"`
}

type Server struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Messaging struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Host       string `yaml:"host"`
	Username   string `yaml:"user"`
	Password   string `yaml:"pass"`
	DbName     string `yaml:"db_name"`
	SkipSchema bool   `yaml:"skip_schema"`
}

func NewConfigFromFile() *Config {
	return readConfig()
}

func readConfig() *Config {
	dir, _ := os.Getwd()
	yamlFile, err := os.Open(fmt.Sprintf("%s/internal/api/configs/configuration.yaml", dir))
	if err != nil {
		log.Fatalln(err)
	}

	defer func(yamlFile *os.File) {
		err := yamlFile.Close()
		if err != nil {
			log.Println("error closing configuration file")
		}
	}(yamlFile)

	conf := &Config{}
	decoder := yaml.NewDecoder(yamlFile)
	err = decoder.Decode(&conf)

	if err != nil {
		log.Fatalln(err)
	}

	return conf
}
