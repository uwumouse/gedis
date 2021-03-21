clean:
	rm -rf ./build
build:
	go build -o build/gserv ./server; go build -o build/gedis ./client