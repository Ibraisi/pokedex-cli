default:
    @just --list

_build:
    go build -o pokedex ./cmd/

run:
    go run ./cmd/

test:
    go test ./... -v

fmt:
    go fmt ./...

vet:
    go vet ./...

check: fmt vet test

clean:
    rm -f pokedex

clean-cache:
    go clean -cache
