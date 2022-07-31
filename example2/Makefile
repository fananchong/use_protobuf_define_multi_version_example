
export PATH := $(shell pwd)/protoc/bin:$(PATH)
PROTOC_VERSION := 21.4
PROTOC_ZIP := protoc-${PROTOC_VERSION}-linux-x86_64.zip
PROTOC_GEN_GO_VERSION := 1.28.1
PROTOC_GEN_GO_VERSION_ZIP := protoc-gen-go.v${PROTOC_GEN_GO_VERSION}.linux.amd64.tar.gz
GOBIN := $(shell go env GOPATH)/bin

proto:
	go mod vendor
	protoc -I=. -I=./vendor --gofast_out=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. examplepb/*.proto
	sed -i 's/_ "."//g' ./examplepb/*.pb.go
	go fmt ./examplepb/*.pb.go

install_protoc:
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}
	unzip -o -d protoc ${PROTOC_ZIP}
	rm -f ${PROTOC_ZIP}
	curl -OL https://github.com/protocolbuffers/protobuf-go/releases/download/v${PROTOC_GEN_GO_VERSION}/${PROTOC_GEN_GO_VERSION_ZIP}
	tar -zxvf ${PROTOC_GEN_GO_VERSION_ZIP} -C protoc/bin
	rm -f ${PROTOC_GEN_GO_VERSION_ZIP}
	go install github.com/gogo/protobuf/protoc-gen-gofast@latest
	mv -f ${GOBIN}/protoc-gen-gofast protoc/bin
	go install github.com/gogo/protobuf/protoc-gen-gogofaster@latest
	mv -f ${GOBIN}/protoc-gen-gogofaster protoc/bin

.PHONY: proto protoc
