#!/bin/sh
GO_ENV=prod CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build life.go
docker build -t denyu95/life .
docker save -o ~/Downloads/life_image.tar  denyu95/life:latest