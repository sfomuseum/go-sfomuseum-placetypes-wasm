GOROOT=$(shell go env GOROOT)
GOMOD=readonly

rebuild:
	@make wasm

wasm:
	GOOS=js GOARCH=wasm go build -mod $(GOMOD) -ldflags="-s -w" -o static/wasm/sfomuseum_placetypes.wasm cmd/placetypes/main.go

example:
	go run cmd/example/main.go -port 8080
