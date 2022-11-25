gen:
	protoc --go_out=.  --go-grpc_out=. --grpc-gateway_out=. proto/*.proto

clean:
	rm pb/*.go

server:
	go run cmd/server/main.go -port 8081

client:
	go run cmd/client/main.go -address 0.0.0.0:8081

.PHONY: gen clean server client test