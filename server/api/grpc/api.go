package grpc

import (
	"context"
	"flag"
	"fmt"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	"github.com/AntonyIS/go-grpc-crud/server/domain"
	repo "github.com/AntonyIS/go-grpc-crud/server/repository/postgres"
)

var (
	port = flag.Int("port", 50051, "server port")
)

type server struct {
	pb.UnimplementedMovieServiceServer
}

func (s *server) CreateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.MovieResponse, error) {
	id := in.GetId()
	name := in.GetName()
	description := in.GetDescription()
	release_date := in.GetReleaseDate()
	image := in.GetImage()

	db, err := repo.NewPostgresRepository()
	if err != nil {
		return nil, nil
	}
	movie := domain.Movie{
		ID:          id,
		Name:        name,
		Description: description,
		ReleaseDate: release_date,
		Image:       image,
	}
	_, err = db.CreateMovie(&movie)
	if err != nil {
		return nil, nil
	}

	return &pb.MovieResponse{Id: id, Name: name, Description: description, ReleaseDate: release_date, Image: image}, nil
}

func (s *server) GetMovie(ctx context.Context, in *pb.MovieID) (*pb.MovieResponse, error) {
	id := in.GetId()
	db, err := repo.NewPostgresRepository()

	if err != nil {
		return nil, nil
	}
	movie, err := db.GetMovie(id)
	if err != nil {
		return nil, nil
	}

	return &pb.MovieResponse{Id: id, Name: movie.ID, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Image: movie.Image}, nil
}

func (s *server) UpdateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.MovieResponse, error) {
	id := in.GetId()
	name := in.GetName()
	description := in.GetDescription()
	release_date := in.GetReleaseDate()
	image := in.GetImage()

	db, err := repo.NewPostgresRepository()
	if err != nil {
		return nil, nil
	}
	movie := domain.Movie{
		ID:          id,
		Name:        name,
		Description: description,
		ReleaseDate: release_date,
		Image:       image,
	}
	_, err = db.UpdateMovie(&movie)
	if err != nil {
		return nil, nil
	}

	return &pb.MovieResponse{Id: id, Name: name, Description: description, ReleaseDate: release_date, Image: image}, nil
}

func (s *server) DeleteMovie(ctx context.Context, in *pb.MovieID) (*pb.Message, error) {
	id := in.GetId()
	db, err := repo.NewPostgresRepository()

	if err != nil {
		return nil, nil
	}
	err = db.DeleteMovie(id)
	if err != nil {
		return nil, nil
	}

	return &pb.Message{Message: fmt.Sprintf("movie with id %s deleted", id)}, nil
}
