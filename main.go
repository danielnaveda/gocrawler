package main

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/danielnaveda/gocrawler/conf"
	"github.com/danielnaveda/gocrawler/files"
	"github.com/danielnaveda/gocrawler/worker"
	yaml "gopkg.in/yaml.v2"
)

func main() {

	var c conf.Conf

	c.GetConf(`conf/craw.yaml`, ioutil.ReadFile, yaml.Unmarshal)

	files.CreateDirIfNotExist(`temp-files`, os.MkdirAll, os.RemoveAll)

	domains := c.Domains

	var wg sync.WaitGroup

	for index := range domains {
		wg.Add(1)
		go worker.CrawlDomain(&c, domains[index]+"/sitemap.xml", &wg)
	}

	wg.Wait()
}
