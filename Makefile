all: generate

fmt:
	go fmt ./...

glide-install:
	glide install --force

logrus-fix:
	rm -fr vendor/github.com/Sirupsen
	find vendor -type f -exec sed -i 's/Sirupsen/sirupsen/g' {} +

generate: generate-proto

generate-proto:
	protoc --gogofaster_out=Mgoogle/protobuf/any.proto=github.com/gogo/protobuf/types,plugins=grpc:. -I. -I$(GOPATH)/src machine.proto


clean-proto:
	rm -fr *pb.go

clean: clean-models

travis: install-deps glide-install logrus-fix generate
	echo "building..."
	go build
