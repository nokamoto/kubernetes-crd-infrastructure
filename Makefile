
all: crd test
	cd crd; make
	# git diff --exit-code

crd:
	mkdir crd
	cd crd; go mod init github.com/nokamoto/kubernetes-crd-infrastructure/crd
	cd crd; kubebuilder init --domain nokamoto.github.com
	cd crd; kubebuilder create api --group infrastructure --version v1 --kind VirtualMachine

api/protobuf/virtual_machine.pb.go: api/protobuf/virtual_machine.proto
	protoc -I api/protobuf api/protobuf/virtual_machine.proto --go_out=plugins=grpc,paths=source_relative:api/protobuf

test:
	cd api; go fmt ./...
	cd api; go test ./...

fmt: api/protobuf/virtual_machine.pb.go
	cd api; go fmt ./...
	cd crd; make fmt

docker-build:
	docker build -t ${IMG} -f Dockerfile .
