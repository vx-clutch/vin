all: build

build:
	@go build -o ./go-out/bin/vin ./cmd/vin/main.go

install:
	@go install ./cmd/vin/main.go
	@mv ~/go/bin/main ~/go/bin/vin

clean:
	@-rm ./go-out/bin/*

.PHONY: all build install clean uninstall
