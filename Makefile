.PHONY: prepare
prepare:
	dep ensure
	cp -R vendor_patches/* vendor

.PHONY: build
build:
	GOOS=js GOARCH=wasm ${BINARY_PATH}go build -tags netgo -o thorchain_wasm_client.wasm ./go
	mkdir -p ./js/dist
	mv thorchain_wasm_client.wasm ./js/dist
