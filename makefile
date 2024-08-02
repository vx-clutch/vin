all: build

build:
	@go build -o ./go-out/bin/vin ./cmd/vin/main.go

install:
	@go install -o vin ./cmd/vin/main.go

clean:
	@-rm ./go-out/bin/*

.PHONY: all build install clean uninstall
