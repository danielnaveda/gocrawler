[![Build Status](https://travis-ci.com/danielnaveda/gocrawler.svg?branch=add-tests)](https://travis-ci.com/danielnaveda/gocrawler)

# Gocrawler

Simple crawler based on the sitemap of a URL

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
