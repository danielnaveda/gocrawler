[![Build Status](https://travis-ci.com/danielnaveda/gocrawler.svg?branch=add-tests)](https://travis-ci.com/danielnaveda/gocrawler)

# Gocrawler

Simple crawler based on the sitemap of a URL

## Run it with Docker

### Without saving files in your machine
```bash
docker pull danielnaveda/gocrawler

docker run -it --rm --name my-running-go-app -v <full path of your custom conf file>:/go/conf/craw.yaml danielnaveda/gocrawler

```

### Saving files in your machine
```bash
docker pull danielnaveda/gocrawler

docker run -it --rm --name my-running-go-app -v <full path of your custom conf file>:/go/conf/craw.yaml -v <full path of the folder where you want to save your files>:/go/temp-files danielnaveda/gocrawler

```


## Quick Installation
Go to your GOPATH and execute:
```bash

# Download project
go get github.com/danielnaveda/gocrawler

# Create your configuration folder and file
mkdir -p conf
cp src/github.com/danielnaveda/gocrawler/conf.dist.yaml ./conf/craw.yaml

# ...<modify ./craw.yaml according to your needs>...

# Run the program
go run src/github.com/danielnaveda/gocrawler/main.go
```
