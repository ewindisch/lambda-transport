#!/bin/bash
go tool cgo -objdir go-build main.go lambda.go
cd node
node-gyp build
