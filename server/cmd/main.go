package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	svr "github.com/AntonyIS/go-grpc-crud/server/api/grpc"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	svr := svr.Server{}
	grpcServer := grpc.NewServer()
	pb.RegisterMovieServiceServer(grpcServer, &svr)
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	fmt.Println("GRPC POSTGRES CRUD")
}
