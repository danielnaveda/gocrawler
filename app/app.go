package app

import (
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/danielnaveda/gocrawler/conf"
	"github.com/danielnaveda/gocrawler/files"
	"github.com/danielnaveda/gocrawler/worker"
	"github.com/olivere/elastic"
)

const (
	folderName  = `temp-files`
	sitemapPath = `sitemap.xml`
)

// Run starts the gocrawler application
func Run() {
	c, err := conf.GetConf()

	var esclient *elastic.Client

	if c.SaveIntoElasticsearch == true {
		esclient, err = elastic.NewClient()
		if err != nil {
			panic(err)
		}

		// Check if index exists
		exists, err := esclient.IndexExists("webpages").Do(context.Background())
		if err != nil {
			panic(err)
		}
		if exists {
			fmt.Println(`Index already exists`)
		} else {
			fmt.Println(`Creating index...`)
			_, err = esclient.CreateIndex("webpages").Do(context.Background())
			if err != nil {
				panic(err)
			}
		}

		// webpage := worker.Webpage{URL: `https://www.somewebsite.com`, Webpage: `<html><head></head><body></body></html>`}
		// _, err = esclient.Index().
		// 	Index("webpages").
		// 	Type("doc").
		// 	BodyJson(webpage).
		// 	Refresh("wait_for").
		// 	Do(context.Background())
		// if err != nil {
		// 	panic(err)
		// }
	}

	if err != nil {
		panic(err.Error())
	}

	files.CreateDirIfNotExist(folderName, os.MkdirAll, os.RemoveAll)

	var wg sync.WaitGroup

	for index := range c.Domains {
		wg.Add(1)
		go worker.CrawlDomain(&c, esclient, c.Domains[index]+"/"+sitemapPath, &wg)
	}

	wg.Wait()
}
