.PHONY: all
all:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop

.PHONY: rebuild
rebuild:
	docker-compose up --build

.PHONY: fmt
fmt:
	goimports -v -w src && go fmt ./src

.PHONY: protoc
protoc:
	protoc -I ./src/pb ./src/pb/*.proto --go_out=plugins=grpc:./src/pb

.PHONY: install-proto
install-proto:
	brew install protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go
