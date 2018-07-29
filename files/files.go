package files

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

type Conf struct {
	Domains                  []string `yaml:"domains"`
	API                      string   `yaml:"api"`
	WorkersPerDomain         int      `yaml:"workers_per_domain"`
	MaxPagesCrawledPerDomain int      `yaml:"max_pages_crawled_per_domain"`
}

func (c *Conf) GetConf() *Conf {

	yamlFile, err := ioutil.ReadFile("conf.yaml")

	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func CreateDirIfNotExist(dir string) {
	os.RemoveAll(dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
