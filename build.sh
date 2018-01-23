#!/bin/bash
#go tool cgo -exportheader lambda.h lambda.go
#-objdir go-build lambda.go
go build -buildmode=c-shared
cd node
node-gyp build
