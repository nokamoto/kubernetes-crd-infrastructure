package main

import (
	"github.com/google/go-cmp/cmp"
	pb "github.com/nokamoto/kubernetes-crd-infrastructure/api/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
)

var foo = &pb.VirtualMachine{Name: "foo"}

func TestVirtualMachineService_Create(t *testing.T) {
	svc := &VirtualMachineService{}
	_, err := svc.Create(nil, foo)
	if err != nil {
		t.Fatal(err)
	}

	_, err = svc.Create(nil, foo)
	if stat := status.Code(err); stat != codes.AlreadyExists {
		t.Fatal(err)
	}
}

func TestVirtualMachineService_Delete(t *testing.T) {
	svc := &VirtualMachineService{}
	_, err := svc.Create(nil, foo)
	if err != nil {
		t.Fatal(err)
	}

	_, err = svc.Delete(nil, foo)
	if err != nil {
		t.Fatal(err)
	}

	_, err = svc.Delete(nil, foo)
	if stat := status.Code(err); stat != codes.NotFound {
		t.Fatal(err)
	}
}

func TestVirtualMachineService_FindByName(t *testing.T) {
	svc := &VirtualMachineService{}
	_, err := svc.Create(nil, foo)
	if err != nil {
		t.Fatal(err)
	}

	res, err := svc.FindByName(nil, foo)
	if err != nil {
		t.Fatal(err)
	}
	if diff := cmp.Diff(res, foo); len(diff) != 0 {
		t.Fatal(diff)
	}

	_, err = svc.Delete(nil, foo)
	if err != nil {
		t.Fatal(err)
	}

	_, err = svc.FindByName(nil, foo)
	if stat := status.Code(err); stat != codes.NotFound {
		t.Fatal(err)
	}
}
