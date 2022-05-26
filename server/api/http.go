package api

import (
	"log"

	"github.com/AntonyIS/go-grpc-crud/server/domain"
	repo "github.com/AntonyIS/go-grpc-crud/server/repository/postgres"
)

func Service() *domain.MovieService {
	db, err := repo.NewPostgresRepository()

	if err != nil {
		log.Fatalf("unable to access service %v", err)
	}
	srv := domain.NewMovieService(db)

	return &srv
}
