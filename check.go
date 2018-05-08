package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Name             string `yaml:"name"`
	BasePatch        string `yaml:"base_patch"`
	Age              int32  `yaml:"age"`
	DatabaseURL      string `yaml:"db_url"`
	DatabasePassword string `yaml:"db_password"`
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func main() {
	var c conf
	c.getConf()

	fmt.Println(c)
}
