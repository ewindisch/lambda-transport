#!/bin/bash
go build -buildmode=c-archive -o node/gorpc.a
cd node
node-gyp configure
node-gyp build
