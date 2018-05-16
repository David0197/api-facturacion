package conf

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type conf struct {
	Version string `yaml:"version"`

	Root string `yaml:"root"`

	Database struct {
		Host     string `yaml:"host"`
		Password string `yaml:"password"`
		Port     string `yaml:"port"`
	}

	CertificatePassword string `yaml:"p12_password"`

	Path struct {
		P12         string `yaml:"p12"`
		Private     string `yaml:"private_key"`
		Certificate string `yaml:"certificates"`
	}
}

func (c *conf) getConf() *conf {

	yamlFile, err := ioutil.ReadFile("./config.yml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

// GetVariable .....
func GetVariable(value string) string {

	var c conf
	c.getConf()

	variables := make(map[string]string)

	variables["version"] = c.Version
	variables["host"] = c.Database.Host
	variables["password"] = c.Database.Password
	variables["port"] = c.Database.Port

	variables["root"] = c.Root

	variables["p12_password"] = c.CertificatePassword

	variables["p12"] = c.Root + c.Path.P12
	variables["private"] = c.Root + c.Path.Private
	variables["certificate"] = c.Root + c.Path.Certificate

	return variables[value]
}
