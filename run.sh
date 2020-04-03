#!/bin/sh
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build life.go
docker build -t denyu95/life .
rm -rf ~/Downloads/life_image.tar
docker save -o ~/Downloads/life_image.tar  denyu95/life:latest