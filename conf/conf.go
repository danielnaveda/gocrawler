package conf

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/crypto/ssh/terminal"
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
	SaveIntoElasticsearch    bool     `yaml:"save_into_elasticsearch"`
	ElasticsearchURL         string   `yaml:"elasticsearch_url"`
	BasicUser                string
	BasicPass                string
	ConfigFilePath           string
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
	var c = Conf{}
	getConfFromDefault(&c)
	getConfFromCLI(&c)
	getConfFromFile(&c)

	return c, nil
}

func getConfFromDefault(c *Conf) {
	c.API = ``
	c.WorkersPerDomain = 10
	c.MaxPagesCrawledPerDomain = -1
	c.SaveIntoFiles = false
	c.BasicUser = ``
	c.BasicPass = ``
	c.ConfigFilePath = ``
	c.SaveIntoElasticsearch = false
	c.ElasticsearchURL = `http://127.0.0.1:9200`
}

func getConfFromFile(c *Conf) {
	if c.ConfigFilePath != `` {
		c.getConfFile(c.ConfigFilePath, ioutil.ReadFile, yaml.Unmarshal)
	}
}

func getConfFromCLI(c *Conf) {
	var myDomainsFromParam domainsFromParam

	confFilePath := flag.String("conf", c.ConfigFilePath, "path of the configuration file")

	apiURL := flag.String("api", c.API, "api url")
	flag.Var(&myDomainsFromParam, "domain", "domain to crawl")
	workersPerDomain := flag.Int("nworkers", c.WorkersPerDomain, "number of worker per domains")
	crawlersPerDomain := flag.Int("ncrawlers", c.MaxPagesCrawledPerDomain, "number of crawlers per domains")
	saveFiles := flag.Bool("savefile", c.SaveIntoFiles, "type false if you do not want to save the files in a folder")

	saveIntoElasticsearch := flag.Bool("saveintoes", c.SaveIntoElasticsearch, "")
	elasticsearchURL := flag.String("esurl", c.ElasticsearchURL, "")

	basicUser := flag.String("basicuser", c.BasicUser, "basic authentication username")
	basicPass := flag.String("basicpass", c.BasicPass, "basic authentication password")

	flag.Parse()

	if *basicUser != "" && *basicPass == "" {
		fmt.Println("User:", *basicUser)
		basicPass = credentials()
	}

	var mydomains []string

	c.API = *apiURL
	c.WorkersPerDomain = *workersPerDomain
	c.MaxPagesCrawledPerDomain = *crawlersPerDomain
	c.SaveIntoFiles = *saveFiles
	c.BasicUser = *basicUser
	c.BasicPass = *basicPass
	c.ConfigFilePath = *confFilePath
	c.SaveIntoElasticsearch = *saveIntoElasticsearch
	c.ElasticsearchURL = *elasticsearchURL

	mydomains = make([]string, len(myDomainsFromParam))

	for index := range myDomainsFromParam {
		mydomains[index] = myDomainsFromParam[index]
	}
	c.Domains = mydomains
}

func credentials() *string {
	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(0)
	fmt.Println(``)
	password := string(bytePassword)
	password = strings.TrimSpace(password)
	return &password
}
