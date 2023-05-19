.PHONY: build

build:
	@mkdir -p bin
	@go build -o bin/app cmd/main.go
	@cp -r templates bin/
	@cp -r static bin/
	@