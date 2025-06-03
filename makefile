all:
	CGO_ENABLED=0 go build -v
start:
	./quin
clean:
	go fmt ./...