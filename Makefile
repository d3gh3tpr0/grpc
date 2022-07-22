gen:
	protoc --go_out=pb --go-grpc_out=pb  proto/*.proto

clean:
	rm pb/*.go

