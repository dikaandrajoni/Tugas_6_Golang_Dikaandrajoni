
package common

import (
	"os"

	parser "Tugas_6_Golang_Dikaandrajoni/Framework/parser"

	log "github.com/Sirupsen/logrus"
)

//Config stores global configuration loaded from json file
type Configuration struct {
	ListenPort string `yaml:"listenPort"`
	RootURL    string `yaml:"rootUrl"`
	Connection struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Password string `yaml:"password"`
		User     string `yaml:"user"`
		Database string `yaml:"database"`
	}
}

var Config Configuration
var logger *log.Entry

func LoadConfigFromFile(fn *string) {
	if err := parser.LoadYAML(fn, &Config); err != nil {
		log.Error("LoadConfigFromFile() - Failed opening config file")
		os.Exit(1)
	}

	log.Info("Loaded configs: ", Config)

}
