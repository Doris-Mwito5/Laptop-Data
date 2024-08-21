gen:
	protoc -I=proto --go_out=. --go-grpc_out=. \ --proto_path=. \ proto/*.proto

generate:
	protoc -I=proto --go_out=. --go-grpc_out=. proto/*proto
clean:
	rm pb/*.go

run:
	go run main.go	

server:
	go run cmd/server/*.go -port 8080

client:
	go run cmd/client/*.go -address 0.0.0.0:8080
