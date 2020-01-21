package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/nokamoto/kubernetes-crd-infrastructure/api/protobuf"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"sync"
)

type VirtualMachineService struct {
	machines []*pb.VirtualMachine
	mutex    sync.Mutex
}

func toJson(machine *pb.VirtualMachine) string {
	bytes, err := json.Marshal(machine)
	if err != nil {
		return fmt.Sprintf("toJson(%v)", err)
	}
	return string(bytes)
}

func (v *VirtualMachineService) Create(ctx context.Context, req *pb.VirtualMachine) (*pb.VirtualMachine, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	log.Printf("Create(%v)", toJson(req))

	for _, machine := range v.machines {
		if machine.GetName() == req.GetName() {
			return nil, status.Error(codes.AlreadyExists, fmt.Sprintf("%s already exists", req.GetName()))
		}
	}

	v.machines = append(v.machines, req)

	return req, nil
}

func (v *VirtualMachineService) Delete(ctx context.Context, req *pb.VirtualMachine) (*pb.VirtualMachine, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	log.Printf("Delete(%v)", toJson(req))

	var machines []*pb.VirtualMachine
	var deleted *pb.VirtualMachine

	for _, machine := range v.machines {
		if machine.GetName() == req.GetName() {
			deleted = machine
		} else {
			machines = append(machines, machine)
		}
	}

	if deleted == nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%s not found", req.GetName()))
	}

	v.machines = machines

	return req, nil
}

func (v *VirtualMachineService) FindByName(ctx context.Context, req *pb.VirtualMachine) (*pb.VirtualMachine, error) {
	v.mutex.Lock()
	defer v.mutex.Unlock()

	log.Printf("FindByName(%v)", toJson(req))

	for _, machine := range v.machines {
		if machine.GetName() == req.GetName() {
			return machine, nil
		}
	}

	return nil, status.Errorf(codes.NotFound, fmt.Sprintf("%s not found", req.GetName()))
}
