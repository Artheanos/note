package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

type Conf struct {
	MongodbURI              string `yaml:"mongodbURI"`
	CookieDurationInSeconds int    `yaml:"cookie_duration_in_seconds"`
}

func GetConfigFile(filePath string) *Conf {
	var result Conf
	yamlFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &result)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &result
}
