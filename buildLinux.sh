#!/usr/bin/env bash
env GOOS=linux GOARCH=amd64 GOARM=7 go build  server_check.go
#env GO111MODULE=off GOPATH=/Users/sam/dev3/pushserver/go GOOS=linux GOARCH=amd64 GOARM=7 go build  cmd/noteserver/noteserver.go

