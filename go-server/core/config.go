package core

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type config struct {
	PartDomain   string `yaml:"part_domain"`
	PublicDomain string `yaml:"public_domain"`
	Server       struct {
		Port uint `yaml:"port"`
	} `yaml:"server"`
	Storage struct {
		FilePath string `yaml:"file_path"`
		DbFile   string `yaml:"db_file"`
	} `yaml:"storage"`
}

var Config config

func loadConfig() {
	configByte, err := os.ReadFile("config.yml")
	if err != nil {
		log.Panicln(err)
	}
	err = yaml.Unmarshal(configByte, &Config)
	if err != nil {
		return
	}
	if Config.Storage.DbFile == "" {
		log.Panicln("storage.db_file未配置")
	}
	if Config.Storage.FilePath == "" {
		log.Panicln("storage.file_path未配置")
	}
	err = os.MkdirAll(Config.Storage.FilePath, 0755)
	if err != nil {
		log.Panicln(err)
	}
}
