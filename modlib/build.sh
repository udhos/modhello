#!/bin/bash

gofmt -s -w *.go
go tool fix *.go
go tool vet .

go test .

