package worker

import (
	"context"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"

	"github.com/danielnaveda/gocrawler/conf"
	"github.com/danielnaveda/gocrawler/sitemap"
	"github.com/olivere/elastic"
)

// Webpage is a simple representation of a webpage
type Webpage struct {
	URL        string `json:"url"`
	StatusCode int    `json:"statuscode"`
}

// CrawlDomain reads the sitemap.xml of a site and fetches all its urls
func CrawlDomain(c *conf.Conf, esclient *elastic.Client, domain string, wg *sync.WaitGroup) {
	defer wg.Done()
	username := c.BasicUser
	passwd := c.BasicPass

	client := &http.Client{}
	req, err := http.NewRequest("GET", domain, nil)
	req.SetBasicAuth(username, passwd)
	resp, err := client.Do(req)

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
		go urlFetchWorker(c, esclient, w, urls, results)
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

func urlFetchWorker(c *conf.Conf, esclient *elastic.Client, id int, jobs <-chan string, results chan<- bool) {
	for url := range jobs {
		fmt.Println("Fetching", url)

		if c.API != "" {
			url = strings.Replace(url, "https://", "", -1)
			url = strings.Replace(url, "http://", "", -1)
		}

		// username := c.BasicUser
		// passwd := c.BasicPass

		client := &http.Client{}
		req, err := http.NewRequest("GET", c.API+url, nil)
		// req.SetBasicAuth(username, passwd)
		resp, err := client.Do(req)

		if err != nil {
			fmt.Println("Error with url " + url)
			results <- true
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)

		if c.SaveIntoFiles == true {
			saveIntoFile(url, body)
		}

		if c.SaveIntoElasticsearch == true {
			webpage := Webpage{URL: url, StatusCode: resp.StatusCode}
			_, err = esclient.Index().
				Index("webpages").
				Type("doc").
				BodyJson(webpage).
				Refresh("wait_for").
				Do(context.Background())
			if err != nil {
				panic(err)
			}
		}

		results <- true
	}
}

func saveIntoFile(filename string, content []byte) error {
	filename = strings.Replace(filename, "/", "..", -1)
	return ioutil.WriteFile("./temp-files/"+filename, content, 0644)
}
