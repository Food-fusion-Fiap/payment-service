.PHONY: default
default: build

all: clean get-deps build test

version := "0.1.0"

build:
	mkdir -p bin
	go build -o bin/payment-service main.go

test: build
	go test -short -coverprofile=bin/cov.out `go list ./... | grep -v vendor/`
	go tool cover -func=bin/coverage.out

clean:
	rm -rf ./bin

sonar: test
	sonar-scanner -Dsonar.projectVersion="$(version)"