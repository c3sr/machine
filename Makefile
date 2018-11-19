all: generate

fmt:
	go fmt ./...


generate: generate-proto

generate-proto:
	protoc --gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,plugins=grpc:. -I. -I$(GOPATH)/src machine.proto


clean-proto:
	rm -fr *pb.go

clean: clean-models

travis: generate
	echo "building..."
	go build
