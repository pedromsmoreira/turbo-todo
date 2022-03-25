package configs

import (
	"embed"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

var config embed.FS

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
		Username string `yaml:"user"`
		Password string `yaml:"pass"`
	} `yaml:"database"`
}

func (conf *Config) ReadConfig() *Config {
	yamlFile, err := os.Open("configuration.yaml")
	if err != nil {
		log.Fatalln(err)
	}

	defer yamlFile.Close()

	decoder := yaml.NewDecoder(yamlFile)
	err = decoder.Decode(&conf)

	if err != nil {
		log.Fatalln(err)
	}

	return conf
}
