package main

import (
	"os"
	"sync"

	"github.com/danielnaveda/gocrawler/conf"
	"github.com/danielnaveda/gocrawler/files"
	"github.com/danielnaveda/gocrawler/worker"
)

func main() {

	var c conf.Conf

	c.GetConf()

	files.CreateDirIfNotExist(`temp-files`, os.MkdirAll, os.RemoveAll)

	domains := c.Domains

	var wg sync.WaitGroup

	for index := range domains {
		wg.Add(1)
		go worker.CrawlDomain(&c, domains[index]+"/sitemap.xml", &wg)
	}

	wg.Wait()
}
