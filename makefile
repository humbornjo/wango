wango: 
	CGO_ENABLED=0 go build -o ./bin/wango

test: 
	go test -v ./...

clean: 
	rm -rf ./bin/*
	rm -rf ./public/*
