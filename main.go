package main

import (
	"gocrawler/conf"
	"gocrawler/files"
	"gocrawler/worker"
	"sync"
)

func main() {
	var c conf.Conf
	c.GetConf()

	files.CreateDirIfNotExist("temp-files")

	domains := c.Domains

	var wg sync.WaitGroup

	for index := range domains {
		wg.Add(1)
		go worker.CrawlDomain(&c, domains[index]+"/sitemap.xml", &wg)
	}

	wg.Wait()
}
