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

.PHONY: proto
proto-ci:
	@sh ./build.sh