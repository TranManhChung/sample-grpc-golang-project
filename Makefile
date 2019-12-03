
genpb:
	protoc -I/usr/local/include -Iidl \
		-I$$GOPATH/src \
		-I$$GOPATH/src/github.com/gogo/protobuf/protobuf \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
		-I$$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway \
		--gogo_out=plugins=grpc,Mgoogle/protobuf/wrappers.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types:grpc-gen/ \
		idl/Service.proto