package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	svr "github.com/AntonyIS/go-grpc-crud/server/api/grpc"
	api "github.com/AntonyIS/go-grpc-crud/server/api/rest"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

var (
	port     = flag.Int("port", 50051, "The server port")
	REST_API = flag.Bool("r", false, "Default API to run")
)

func main() {
	flag.Parse()

	switch *REST_API {
	case true:
		RESTServer()
	default:
		GRPCServer()
	}

}

func GRPCServer() {
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

func RESTServer() {
	handler := api.NewHandler()
	r := gin.Default()

	r.POST("/api/v1/movies", handler.CreateMovie)
	r.GET("/api/v1/movies", handler.GetMovies)
	r.GET("/api/v1/movies/:id", handler.GetMovie)
	r.PUT("/api/v1/movies", handler.UpdateMovie)
	r.DELETE("/api/v1/movies/:id", handler.DeleteMovie)
	r.Run((":5000"))
}
