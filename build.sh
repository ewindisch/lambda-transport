#!/bin/bash
#go tool cgo -exportheader lambda.h lambda.go
#-objdir go-build lambda.go
go build -buildmode=c-shared -o node/gorpc.so
cd node
node-gyp configure
node-gyp build
