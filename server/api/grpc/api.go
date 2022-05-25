package grpc

import (
	"context"
	"fmt"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	"github.com/AntonyIS/go-grpc-crud/server/domain"
	repo "github.com/AntonyIS/go-grpc-crud/server/repository/postgres"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMovieServiceServer
}

func (s *Server) CreateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.MovieResponse, error) {

	db, err := repo.NewPostgresRepository()
	if err != nil {
		return nil, nil
	}
	srv := domain.NewMovieService(db)
	movie := domain.Movie{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		ReleaseDate: in.GetReleaseDate(),
		Image:       in.GetImage(),
	}
	r, err := srv.CreateMovie(&movie)

	if err != nil {
		return nil, nil
	}

	return &pb.MovieResponse{Id: r.ID, Name: r.Name, Description: r.Description, ReleaseDate: r.ReleaseDate, Image: r.Image}, nil
}

func (s *Server) GetMovie(ctx context.Context, in *pb.MovieID) (*pb.MovieResponse, error) {

	id := in.GetId()
	db, err := repo.NewPostgresRepository()

	if err != nil {
		return nil, nil
	}
	movie, err := db.GetMovie(id)

	if err != nil {
		return &pb.MovieResponse{}, status.Error(404, "movie not found")
	}

	return &pb.MovieResponse{Id: id, Name: movie.ID, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Image: movie.Image}, nil
}
func (s *Server) GetMovies(ctx context.Context, in *pb.EmptyRequest) (*pb.MovieListResponse, error) {

	db, err := repo.NewPostgresRepository()

	if err != nil {
		return nil, nil
	}
	movies, err := db.GetMovies()

	if err != nil {
		return &pb.MovieListResponse{}, status.Error(404, "movies not found")
	}

	return &pb.MovieListResponse{Movies: movies}, nil
}

func (s *Server) UpdateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.MovieResponse, error) {
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

func (s *Server) DeleteMovie(ctx context.Context, in *pb.MovieID) (*pb.Message, error) {
	id := in.GetId()
	db, err := repo.NewPostgresRepository()

	if err != nil {
		return nil, nil
	}
	err = db.DeleteMovie(id)
	if err != nil {
		return nil, nil
	}

	return &pb.Message{Message: fmt.Sprintf("movie with id %s deleted successuly", id)}, nil
}
