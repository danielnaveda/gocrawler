package main

import (
	"os"
	"sync"

	"github.com/danielnaveda/gocrawler/conf"
	"github.com/danielnaveda/gocrawler/files"
	"github.com/danielnaveda/gocrawler/worker"
)

// type OsInterface interface {
// 	RemoveAll(string)
// 	Stat(string) (os.FileInfo, error)
// 	IsNotExist(error) bool
// 	MkdirAll(string, int) error
// }
type workaroundObject struct{}

func (w *workaroundObject) RemoveAll(dir string) {
	os.RemoveAll(dir)
}

func (w *workaroundObject) Stat(dir string) (os.FileInfo, error) {
	return os.Stat(dir)
}

func (w *workaroundObject) IsNotExist(err error) bool {
	return os.IsNotExist(err)
}

func (w *workaroundObject) MkdirAll(dir string, fm os.FileMode) error {
	return os.MkdirAll(dir, fm)
}

func main() {

	myobject := workaroundObject{}

	var c conf.Conf

	c.GetConf()

	files.CreateDirIfNotExist("temp-files", &myobject)

	domains := c.Domains

	var wg sync.WaitGroup

	for index := range domains {
		wg.Add(1)
		go worker.CrawlDomain(&c, domains[index]+"/sitemap.xml", &wg)
	}

	wg.Wait()
}
