#!/bin/bash

gofmt -s -w lib/*.go
go tool fix lib/*.go
go tool vet ./lib
go test ./lib

