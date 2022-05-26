package grpc

import (
	"context"

	pb "github.com/AntonyIS/go-grpc-crud/proto"
	srv "github.com/AntonyIS/go-grpc-crud/server/api"
	"github.com/AntonyIS/go-grpc-crud/server/domain"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedMovieServiceServer
}

func (s *Server) CreateMovie(ctx context.Context, in *pb.MovieRequest) (*pb.MovieResponse, error) {

	srv := *srv.Service()
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
	srv := *srv.Service()
	movie, err := srv.GetMovie(id)

	if err != nil {
		return &pb.MovieResponse{}, status.Error(404, "movie not found")
	}

	return &pb.MovieResponse{Id: id, Name: movie.Name, Description: movie.Description, ReleaseDate: movie.ReleaseDate, Image: movie.Image}, nil
}
func (s *Server) GetMovies(ctx context.Context, in *pb.EmptyRequest) (*pb.MovieListResponse, error) {

	srv := *srv.Service()
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

	srv := *srv.Service()
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
	srv := *srv.Service()
	err := srv.DeleteMovie(id)
	if err != nil {
		return nil, nil
	}

	return &pb.Message{Message: "movie with deleted successuly"}, nil
}
