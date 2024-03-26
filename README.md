# cache2go

[![Latest Release](https://img.shields.io/github/release/muesli/cache2go.svg)](https://github.com/muesli/cache2go/releases)
[![Build Status](https://github.com/muesli/cache2go/workflows/build/badge.svg)](https://github.com/muesli/cache2go/actions)
[![Coverage Status](https://coveralls.io/repos/github/muesli/cache2go/badge.svg?branch=master)](https://coveralls.io/github/muesli/cache2go?branch=master)
[![Go ReportCard](https://goreportcard.com/badge/muesli/cache2go)](https://goreportcard.com/report/muesli/cache2go)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/muesli/cache2go)

Concurrency-safe golang caching library with expiration capabilities.

This is a personal Go learning project, on the basis of the original project added the function of interacting ,added an Update function, optimized the background garbage cleaning function adn added a function that returns all elements in the table, so that the user can more conveniently experience the functions of adding, deleting, querying and so on. The project source address is <https://github.com/muesli/cache2go>, and we would like to thank the original author for his contribution.

## Installation

Make sure you have a working Go environment (Go 1.2 or higher is required).
See the [install instructions](https://golang.org/doc/install.html).

To install cache2go, simply run:

    go get github.com/bhluo/cache2go

To compile it from source:

    cd $GOPATH/src/github.com/bhluo/cache2go
    go get -u -v
    go build && go test -v


To run the code, go to examples/mycachedapp/ and run:

    go run mycachedapp.go
