
export PATH := $(shell pwd)/protoc/bin:$(PATH)
PROTOC_VERSION := 21.4
PROTOC_ZIP := protoc-${PROTOC_VERSION}-linux-x86_64.zip
PROTOC_GEN_GO_VERSION := 1.28.1
PROTOC_GEN_GO_VERSION_ZIP := protoc-gen-go.v${PROTOC_GEN_GO_VERSION}.linux.amd64.tar.gz

proto:
	protoc -I=. -I=./protoc/include --go_out=. proto/*.proto

install_protoc:
	curl -OL https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}
	unzip -f -d protoc ${PROTOC_ZIP}
	rm -f ${PROTOC_ZIP}
	curl -OL https://github.com/protocolbuffers/protobuf-go/releases/download/v${PROTOC_GEN_GO_VERSION}/${PROTOC_GEN_GO_VERSION_ZIP}
	tar -zxvf ${PROTOC_GEN_GO_VERSION_ZIP} -C protoc/bin
	rm -f ${PROTOC_GEN_GO_VERSION_ZIP}

.PHONY: proto protoc