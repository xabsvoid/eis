include platform.mk

build-eis:
	go build -o $(LOCAL_BIN) ./cmd/eis

build: build-eis
