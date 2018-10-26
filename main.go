package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"

	"github.com/danielnaveda/gocrawler/conf"
	"github.com/danielnaveda/gocrawler/files"
	"github.com/danielnaveda/gocrawler/worker"
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

func main() {
	var myDomainsFromParam domainsFromParam

	confFileFlag := flag.String("conf", "", "path of the configuration file")

	apiFlag := flag.String("api", "", "api url")
	myUserName := flag.String("user", "", "Basic auth username")
	myPassword := flag.String("pass", "", "Basic auth password")

	flag.Var(&myDomainsFromParam, "domain", "domain to crawl")

	workersPerDomainFlag := flag.Int("nworkers", 100, "domains separated by commas")
	crawlersPerDomainFlag := flag.Int("ncrawlers", -1, "domains separated by commas")
	saveFilesFlag := flag.Bool("savefile", true, "domains separated by commas")

	flag.Parse()

	fmt.Println("User:", *myUserName)

	if *myUserName != "" && *myPassword == "" {
		myPassword := credentials()
		fmt.Println("Password typed:", myPassword)
	} else {
		fmt.Println("\nPassword from cli:", *myPassword)
	}

	var mydomains []string

	var c = conf.Conf{
		API:                      *apiFlag,
		WorkersPerDomain:         *workersPerDomainFlag,
		MaxPagesCrawledPerDomain: *crawlersPerDomainFlag,
		SaveIntoFiles:            *saveFilesFlag,
	}

	c.GetConf(*confFileFlag, ioutil.ReadFile, yaml.Unmarshal)

	mydomains = make([]string, len(myDomainsFromParam))

	for index := range myDomainsFromParam {
		mydomains[index] = myDomainsFromParam[index]
	}
	c.Domains = mydomains

	files.CreateDirIfNotExist(`temp-files`, os.MkdirAll, os.RemoveAll)

	domains := c.Domains

	var wg sync.WaitGroup

	for index := range domains {
		wg.Add(1)
		go worker.CrawlDomain(&c, domains[index]+"/sitemap.xml", &wg)
	}

	wg.Wait()
}

func credentials() string {
	fmt.Print("Enter Password: ")
	bytePassword, _ := terminal.ReadPassword(0)
	password := string(bytePassword)
	return strings.TrimSpace(password)
}
