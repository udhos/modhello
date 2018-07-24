#!/bin/bash

gofmt -s -w ./*/*.go
go tool fix ./*/*.go
go tool vet ./modlib
go tool vet ./modapp

go test ./modlib
go test ./modapp
go install ./modapp

