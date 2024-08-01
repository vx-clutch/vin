all: build

build:
	@go build -o ./go-out/bin/vin ./cmd/vin/main.go

install: build
	@sudo mv ./go-out/bin/vin /usr/local/bin/vin

uninstall: @-clean
	@-sudo rm /usr/local/bin/vin

clean:
	@-rm ./go-out/bin/*

.PHONY: all build install clean uninstall
