package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	yaml "gopkg.in/yaml.v2"
)

type conf struct {
	Domains                  []string `yaml:"domains"`
	API                      string   `yaml:"api"`
	WorkersPerDomain         int      `yaml:"workers_per_domain"`
	MaxPagesCrawledPerDomain int      `yaml:"max_pages_crawled_per_domain"`
}

func (c *conf) getConf() *conf {

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

// Urlset represents the xml schema of the sitemap
type Urlset struct {
	URL []struct {
		Loc        string `xml:"loc"`
		Lastmod    string `xml:"lastmod"`
		Priority   string `xml:"priority"`
		Changefreq string `xml:"changefreq"`
	} `xml:"url"`
}

func main() {
	var c conf
	c.getConf()

	createDirIfNotExist("files")

	domains := c.Domains

	var wg sync.WaitGroup

	for index := range domains {
		wg.Add(1)
		go crawlDomain(&c, domains[index]+"/sitemap.xml", &wg)
	}

	wg.Wait()
}

func crawlDomain(c *conf, domain string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(domain)
	if err != nil {
		fmt.Println("Error with domain " + domain)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var mysitemap Urlset
	xml.Unmarshal(body, &mysitemap)
	// fmt.Println(mysitemap)

	urls := make(chan string, 20000)
	results := make(chan bool, 20000)

	for w := 1; w <= c.WorkersPerDomain; w++ {
		go urlFetchWorker(c, w, urls, results)
	}

	counter := 0

	for index := range mysitemap.URL {
		fmt.Println(mysitemap.URL[index].Loc)
		urls <- mysitemap.URL[index].Loc
		counter++
		if index == (c.MaxPagesCrawledPerDomain - 1) {
			break
		}
	}

	close(urls)

	for a := 1; a <= counter; a++ {
		<-results
	}
}

func stringModifier(originalURL string) string {
	return originalURL
}

func urlFetchWorker(c *conf, id int, jobs <-chan string, results chan<- bool) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)

		if c.API != "" {
			j = strings.Replace(j, "https://", "", -1)
			j = strings.Replace(j, "http://", "", -1)
		}

		resp2, err2 := http.Get(c.API + j)
		if err2 != nil {
			fmt.Println("Error with url " + j)
			results <- true
			return
		}

		body2, _ := ioutil.ReadAll(resp2.Body)

		d1 := []byte("testing\n")
		d1 = body2
		j = strings.Replace(j, "/", "|", -1)
		// TODO: place this into a folder
		err := ioutil.WriteFile("./files/worker_"+strconv.Itoa(id)+"_"+string(j), d1, 0644)
		fmt.Println(err)

		fmt.Println("worker", id, "finished job", j)
		results <- true
	}
}

func createDirIfNotExist(dir string) {
	os.RemoveAll(dir)

	if _, err := os.Stat(dir); os.IsNotExist(err) {
		err = os.MkdirAll(dir, 0755)
		if err != nil {
			panic(err)
		}
	}
}
