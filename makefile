wango: 
	CGO_ENABLED=0 go build -o ./bin/wango

release: 
	CGO_ENABLED=0 go build -o ./bin/wango -ldflags "-s -w"

test: 
	go test -v ./...

clean: 
	rm -rf ./bin/*
	rm -rf ./*.png
