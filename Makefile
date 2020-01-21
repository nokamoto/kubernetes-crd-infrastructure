
all: crd
	cd crd; make

crd:
	mkdir crd
	cd crd; go mod init github.com/nokamoto/kubernetes-crd-infrastructure/crd
	cd crd; kubebuilder init --domain nokamoto.github.com
	cd crd; kubebuilder create api --group infrastructure --version v1 --kind VirtualMachine
