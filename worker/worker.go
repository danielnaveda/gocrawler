package worker

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/danielnaveda/gocrawler/conf"
	"github.com/danielnaveda/gocrawler/sitemap"
)

func CrawlDomain(c *conf.Conf, domain string, wg *sync.WaitGroup) {
	defer wg.Done()
	resp, err := http.Get(domain)
	if err != nil {
		fmt.Println("Error with domain " + domain)
		return
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var mysitemap sitemap.Urlset
	xml.Unmarshal(body, &mysitemap)

	urls := make(chan string, 20000)
	results := make(chan bool, 20000)

	for w := 1; w <= c.WorkersPerDomain; w++ {
		go urlFetchWorker(c, w, urls, results)
	}

	counter := 0

	for index := range mysitemap.URL {
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

func urlFetchWorker(c *conf.Conf, id int, jobs <-chan string, results chan<- bool) {
	for url := range jobs {
		fmt.Println("Fetching", url)

		if c.API != "" {
			url = strings.Replace(url, "https://", "", -1)
			url = strings.Replace(url, "http://", "", -1)
		}

		resp, err := http.Get(c.API + url)
		if err != nil {
			fmt.Println("Error with url " + url)
			results <- true
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)

		if c.SaveIntoFiles == true {
			saveIntoFile(url, body)
		}

		results <- true
	}
}

func saveIntoFile(filename string, content []byte) error {
	filename = strings.Replace(filename, "/", "..", -1)
	return ioutil.WriteFile("./temp-files/"+filename, content, 0644)
}
