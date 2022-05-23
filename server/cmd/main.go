package main

import (
	"context"
	"log"
	"net"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	"github.com/AntonyIS/go-grpc-crud/server/internals/domain"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type CarServer struct {
	pb.UnimplementedCarServiceServer
	srv domain.CarService
}

func (cs *CarServer) CreateCar(ctx context.Context, in *pb.CarRequest) (*pb.CarResponse, error) {
	c := &domain.Car{
		Name:        in.GetName(),
		Description: in.GetDescription(),
	}
	car, err := cs.srv.CreateCar(c)
	if err != nil {
		log.Println("Error adding new car")
	}

	return &pb.CarResponse{Name: in.GetName(), Description: in.GetDescription(), Id: car.ID}, nil

}
func (cs *CarServer) GetCar(ctx context.Context, in *pb.CarID) (*pb.CarResponse, error) {

	car, err := cs.srv.GetCar(in.GetId())
	if err != nil {
		log.Println("Error adding new car")
	}

	return &pb.CarResponse{Name: car.ID, Description: car.Name, Id: car.ID}, nil

}
func (cs *CarServer) UpdateCar(ctx context.Context, in *pb.CarRequest) (*pb.CarResponse, error) {
	c := &domain.Car{
		Name:        in.GetName(),
		Description: in.GetDescription(),
	}
	car, err := cs.srv.UpdateCar(c)
	if err != nil {
		log.Println("Error adding new car")
	}

	return &pb.CarResponse{Name: car.ID, Description: car.Name, Id: car.ID}, nil

}
func (cs *CarServer) DeleteCar(ctx context.Context, in *pb.CarID) (*pb.CarMessageResponse, error) {
	err := cs.srv.DeleteCar(in.GetId())
	if err != nil {
		log.Println("Error adding new car")
	}

	return &pb.CarMessageResponse{Message: "Car deleted successully"}, nil
}

func main() {
	lis, err := net.Listen("tpc", port)
	if err != nil {
		log.Fatalf("failed to listen : %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterCarServiceServer(s, &CarServer{})

	log.Printf("server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
