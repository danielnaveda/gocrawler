package conf

import (
	"flag"
	"io/ioutil"

	yaml "gopkg.in/yaml.v2"
)

type domainsFromParam []string

func (d *domainsFromParam) String() string {
	return "dummy implementation"
}

func (d *domainsFromParam) Set(value string) error {
	*d = append(*d, value)
	return nil
}

// Conf is the configuration structure
type Conf struct {
	Domains                  []string `yaml:"domains"`
	API                      string   `yaml:"api"`
	WorkersPerDomain         int      `yaml:"workers_per_domain"`
	MaxPagesCrawledPerDomain int      `yaml:"max_pages_crawled_per_domain"`
	SaveIntoFiles            bool     `yaml:"save_into_files"`
	BasicUser                string
	BasicPass                string
}

// getConfFile reads the configuration file
func (c *Conf) getConfFile(fileName string, reader func(string) ([]byte, error), unmarshaller func([]byte, interface{}) error) (*Conf, error) {
	yamlFile, err := reader(fileName)

	if err != nil {
		return nil, err
	}

	err = unmarshaller(yamlFile, c)

	return c, err
}

// GetConf returns a configuration struct based on the user's requirements
func GetConf() (Conf, error) {
	var myDomainsFromParam domainsFromParam

	confFilePath := flag.String("conf", "", "path of the configuration file")

	apiURL := flag.String("api", "", "api url")
	flag.Var(&myDomainsFromParam, "domain", "domain to crawl")
	workersPerDomain := flag.Int("nworkers", 100, "number of worker per domains")
	crawlersPerDomain := flag.Int("ncrawlers", -1, "number of crawlers per domains")
	saveFiles := flag.Bool("savefile", true, "type false if you do not want to save the files in a folder")

	basicUser := flag.String("basicuser", "", "basic authentication username")
	basicPass := flag.String("basicpass", "", "basic authentication password")

	flag.Parse()

	var mydomains []string

	var c = Conf{
		API:                      *apiURL,
		WorkersPerDomain:         *workersPerDomain,
		MaxPagesCrawledPerDomain: *crawlersPerDomain,
		SaveIntoFiles:            *saveFiles,
		BasicUser:                *basicUser,
		BasicPass:                *basicPass,
	}

	c.getConfFile(*confFilePath, ioutil.ReadFile, yaml.Unmarshal)

	mydomains = make([]string, len(myDomainsFromParam))

	for index := range myDomainsFromParam {
		mydomains[index] = myDomainsFromParam[index]
	}
	c.Domains = mydomains

	return c, nil
}
