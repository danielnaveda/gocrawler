# Build binary
FROM golang:alpine AS build-env
RUN apk update && \
    apk add --no-cache git
ENV GOPATH=/go
RUN go get -v github.com/danielnaveda/gocrawler

# Build runtime
FROM alpine
WORKDIR /app
RUN apk update && \
    apk add --no-cache curl
COPY --from=build-env /go/bin/gocrawler /app/
ENTRYPOINT [ "/app/gocrawler" ]
