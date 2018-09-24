FROM golang

RUN go get github.com/danielnaveda/gocrawler
RUN mkdir conf

ENTRYPOINT [ "/go/bin/gocrawler" ]
