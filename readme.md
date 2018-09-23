[![Build Status](https://travis-ci.com/danielnaveda/gocrawler.svg?branch=add-tests)](https://travis-ci.com/danielnaveda/gocrawler)

# Gocrawler

Simple crawler based on the sitemap of a URL

## Run it with Docker

### Quick run
```bash
docker run danielnaveda/gocrawler -domain "https://www.a-domain-with-a-standard-sitemap.com"
```

### Keep the crawled files in your machine
```bash
docker run -v <your local full path>:/go/temp-files danielnaveda/gocrawler -domain "https://www.a-domain-with-a-standard-sitemap.com"
```
