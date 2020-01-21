package main

import (
	"flag"
	"fmt"
	pb "github.com/nokamoto/kubernetes-crd-infrastructure/api/protobuf"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	port = flag.Int("port", 9090, "gRPC サーバーポート")
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("listen tcp port (%d) - %v", *port, err)
	}

	var opts []grpc.ServerOption
	server := grpc.NewServer(opts...)
	pb.RegisterVirtualMachineServiceServer(server, &VirtualMachineService{})

	log.Printf("ready to serve %d", *port)

	err = server.Serve(lis)
	if err != nil {
		log.Fatalf("serve %v - %v", lis, err)
	}
}
