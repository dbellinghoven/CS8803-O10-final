#!/bin/sh

set -e

wget "https://www.fhfa.gov/DataTools/Downloads/Documents/Enterprise-PUDB/National-File-A/2018_SFNationalFileA2018.zip"
unzip 2018_SFNationalFileA2018.zip
docker run -v `pwd`:/go/src/app golang:1.14.6-alpine3.12 go run /go/src/app/cleanup.go
