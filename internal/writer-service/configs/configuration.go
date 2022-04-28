package configs

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"server"`
	Messaging struct {
		Host string `yaml:"host"`
		Port string `yaml:"port"`
	} `yaml:"messaging"`
	Database struct {
		Username   string `yaml:"user"`
		Password   string `yaml:"pass"`
		SkipSchema bool   `yaml:"skip_schema"`
	} `yaml:"database"`
}

func (conf *Config) ReadConfig() *Config {
	yamlFile, err := os.Open("configuration.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	defer func(yamlFile *os.File) {
		err := yamlFile.Close()
		if err != nil {
			log.Println("error closing configuration file")
		}
	}(yamlFile)

	decoder := yaml.NewDecoder(yamlFile)
	err = decoder.Decode(&conf)

	if err != nil {
		log.Fatalln(err)
	}

	return conf
}
