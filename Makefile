build:
	go build -o build/

dev:
	go run . -local -config config.json

test:
	go test -v ./... -bench ./... -benchmem

# https://nixpacks.com/docs/providers/go
nixpacks-run:
	./out -config config.json