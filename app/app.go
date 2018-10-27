package app

import (
	"os"
	"sync"

	"github.com/danielnaveda/gocrawler/conf"
	"github.com/danielnaveda/gocrawler/files"
	"github.com/danielnaveda/gocrawler/worker"
)

// Run starts the gocrawler application
func Run() {
	c, err := conf.GetConf()

	if err != nil {
		panic(err.Error())
	}

	files.CreateDirIfNotExist(`temp-files`, os.MkdirAll, os.RemoveAll)

	var wg sync.WaitGroup

	for index := range c.Domains {
		wg.Add(1)
		go worker.CrawlDomain(&c, c.Domains[index]+"/sitemap.xml", &wg)
	}

	wg.Wait()
}
