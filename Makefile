.PHONY: build
build:
	@GOOS=linux GOARCH=amd64 go build -o bin/app cmd/main.go

.PHONY: build-image
build-image: build
	@docker build -t simple-callback .


.PHONY: proto
proto:
	@protoc -I/usr/local/protobuffer/include -I. --go_out=. --go-grpc_out=. ./pb/*.proto
	@mv simple-callback/pbs/*.go pb/
	@rm -rf simple-callback

.PHONY:install-go-plugins
install-go-plugins:
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

.PHONY: build-proto
build-proto: install-go-plugins
	@bash build.sh build

.PHONY:
tidy:
	@go mod tidy

success:
	@bash build.sh success

fail:
	@bash build.sh fail