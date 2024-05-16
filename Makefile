build :
	go build -o bin/price-fetcher 

run : build
	./bin/price-fetcher

proto :
	protoc --go_out=. --go_out=path=source=relative \
		--go-grpc_out=. --go-grpc_out=path=source=relative \
		proto/service.proto

.PHONY: proto
