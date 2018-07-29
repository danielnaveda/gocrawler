package worker

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
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

func urlFetchWorker(c *conf.Conf, id int, jobs <-chan string, results chan<- bool) {
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
		err := ioutil.WriteFile("./temp-files/worker_"+strconv.Itoa(id)+"_"+string(j), d1, 0644)
		fmt.Println(err)

		fmt.Println("worker", id, "finished job", j)
		results <- true
	}
}
