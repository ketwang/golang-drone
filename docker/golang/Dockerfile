FROM golang:1.12
MAINTAINER ketwang
COPY date /usr/bin/
COPY golangci-lint /usr/local/bin/gotools/
COPY staticcheck  /usr/local/bin/gotools/
ENV  PATH /usr/local/bin/gotools/:$PATH