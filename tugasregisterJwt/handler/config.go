package handler

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type sqlDb struct {
	Sqlserver   string `yaml:"sqlserver"`
	Sqlport     int    `yaml:"sqlport"`
	SqldbName   string `yaml:"sqldbName"`
	Sqluser     string `yaml:"sqluser"`
	Sqlpassword string `yaml:"sqlpassword"`
}

var configYaml = "config/tugasregisterjwt.yaml"

type configuration struct {
	// Raw file data to avoid re-reading of configuration file
	// It's reset after config is parsed
	ConnectionString sqlDb  `yaml:"sqldatabase"`
	SecretKey        string `yaml:"lusia"`
}

var config = configuration{}

func ParseConfig() error {
	yamlFile, err := readConfigFile()
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return err
	}

	return nil
}

func readConfigFile() ([]byte, error) {
	os.Chdir(".")
	A, _ := os.Getwd()
	fmt.Print(A)
	d, err := ioutil.ReadFile(configYaml)
	if err != nil {
		return nil, err
	}
	return d, nil
}
