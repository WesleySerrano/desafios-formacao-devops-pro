#!/bin/sh
docker stop docker-go-local
docker rm docker-go-local
docker build --label docker-go -t docker-go:0.0.1 -t docker-go:latest  -f dockerfile-go .
docker run -d -l docker-go-local --name docker-go-local -p 8181:8181 docker-go:0.0.1
docker image prune