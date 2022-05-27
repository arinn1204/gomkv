build:
	@go build ./cmd/gomkv

test:
	@go test ./...

lint:
	@go fmt ./... && golint -set_exit_status internal/... cmd/... pkg/...

clean:
	rm gomkv
