build: build-node build-python

build-node:
	go build -buildmode=c-archive -o node/gorpc.a

build-python:
	go build -buildmode=c-shared -o python/gorpc.so

clean: clean-node clean-python

clean-node:
	rm -rf node/build
	rm -f node/gorpc.a
	rm -f node/gorpc.h

clean-python:
	rm -rf python/__pycache__
	rm -f python/gorpc.h
	rm -f python/gorpc.so

node-gyp:
	cd node && \
		node-gyp configure && \
		node-gyp build

node: clean-node build-node node-gyp

python: clean-python build-python
