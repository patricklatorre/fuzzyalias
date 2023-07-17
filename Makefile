build:
	go build -o build/

dev:
	go run . -local -config config.json

test:
	go test -v ./... -bench ./... -benchmem