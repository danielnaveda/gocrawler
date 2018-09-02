FROM golang

RUN go get github.com/danielnaveda/gocrawler
RUN mkdir conf

CMD /go/bin/gocrawler

