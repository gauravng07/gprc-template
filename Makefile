SERVICE ?= grpc-template
BUILDENV :=
BUILDENV += GO111MODULE=on CGO_ENABLED=0

protos:
	GO111MODULE=on ./build/proto.sh