#!/usr/bin/env bash

make clean
make test-xml
make test-coverage
go test -short -coverprofile=bin/cov.out `go list ./... | grep -v vendor/`
make
