
all: crd api/protobuf/virtual_machine.pb.go
	cd crd; make

crd:
	mkdir crd
	cd crd; go mod init github.com/nokamoto/kubernetes-crd-infrastructure/crd
	cd crd; kubebuilder init --domain nokamoto.github.com
	cd crd; kubebuilder create api --group infrastructure --version v1 --kind VirtualMachine

api/protobuf/virtual_machine.pb.go: api/protobuf/virtual_machine.proto
	protoc -I api/protobuf api/protobuf/virtual_machine.proto --go_out=plugins=grpc,paths=source_relative:api/protobuf
