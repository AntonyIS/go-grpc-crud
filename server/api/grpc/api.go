package grpc

import (
	"context"
	"fmt"
	"log"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	"github.com/AntonyIS/go-grpc-crud/server/domain"
	repo "github.com/AntonyIS/go-grpc-crud/server/repository/postgres"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMovieServiceServer
}

func (s *Server) CreateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.MovieResponse, error) {

	srv := *service()
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
	srv := *service()
	movie, err := srv.GetMovie(id)

	if err != nil {
		return &pb.MovieResponse{}, status.Error(404, "movie not found")
	}

	return &pb.MovieResponse{Id: id, Name: movie.Name, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Image: movie.Image}, nil
}
func (s *Server) GetMovies(ctx context.Context, in *pb.EmptyRequest) (*pb.MovieListResponse, error) {

	srv := *service()
	movies, err := srv.GetMovies()
	if err != nil {
		return &pb.MovieListResponse{}, status.Error(404, "movies not found")
	}
	var movieReponse = []*pb.MovieResponse{}
	for _, movie := range *movies {
		movieReponse = append(movieReponse, &pb.MovieResponse{Id: movie.ID, Name: movie.Name, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Image: movie.Image})
	}
	return &pb.MovieListResponse{Movies: movieReponse}, nil
}

func (s *Server) UpdateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.MovieResponse, error) {

	srv := *service()
	movie := domain.Movie{
		ID:          in.GetId(),
		Name:        in.GetName(),
		Description: in.GetDescription(),
		ReleaseDate: in.GetReleaseDate(),
		Image:       in.GetImage(),
	}
	_, err := srv.UpdateMovie(&movie)
	if err != nil {
		return nil, nil
	}

	return &pb.MovieResponse{Id: movie.ID, Name: movie.Name, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Image: movie.Image}, nil
}

func (s *Server) DeleteMovie(ctx context.Context, in *pb.MovieID) (*pb.Message, error) {
	id := in.GetId()
	srv := *service()
	err := srv.DeleteMovie(id)
	if err != nil {
		return nil, nil
	}

	return &pb.Message{Message: fmt.Sprintf("movie with id %s deleted successuly", id)}, nil
}

func service() *domain.MovieService {
	db, err := repo.NewPostgresRepository()

	if err != nil {
		log.Fatalf("unable to access service %v", err)
	}
	srv := domain.NewMovieService(db)

	return &srv
}
