#!/bin/sh

GOPATH=$(go env GOPATH)
echo $GOPATH

[ -d $GOPATH/bin/protoc-gen-gogoslick ] || go get github.com/gogo/protobuf/protoc-gen-gogoslick
[ -d $GOPATH/bin/protoc-gen-gofast ] || go get github.com/gogo/protobuf/protoc-gen-gofast
[ -d $GOPATH/bin/protoc-min-version ] || go get github.com/gogo/protobuf/protoc-min-version
[ -d $GOPATH/src/google.golang.org/grpc ] || go get -u google.golang.org/grpc
[ -d $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway ] || go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
[ -d $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger ] || go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger

package="api"
protoPath="./proto"

rm -rf pkg/pb

mkdir -p pkg/pb/generated/$package
[ -d $goout ] || mkdir -p $goout
overrides=Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types

protoc-min-version \
  --version="3.0.0" \
  -I $protoPath \
  -I $GOPATH/src \
  -I $GOPATH/src/github.com/gogo/protobuf/protobuf \
  -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --gogoslick_out=$overrides,plugins=grpc,import_path=$package:pkg/pb/generated/$package \
  $protoPath/*.proto
