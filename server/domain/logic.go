package domain

import (
	"github.com/teris-io/shortid"
)

type movieService struct {
	movieRepo MovieRepository
}

func NewMovieService(movieRepo MovieRepository) MovieService {
	return &movieService{
		movieRepo,
	}
}

func (ms movieService) CreateMovie(movie *Movie) (*Movie, error) {

	movie.ID = shortid.MustGenerate()
	return ms.movieRepo.CreateMovie(movie)
}

func (ms movieService) GetMovie(id string) (*Movie, error) {
	return ms.movieRepo.GetMovie(id)
}

func (ms movieService) GetMovies() (*[]Movie, error) {
	return ms.movieRepo.GetMovies()
}

func (ms movieService) UpdateMovie(movie *Movie) (*Movie, error) {
	return ms.movieRepo.UpdateMovie(movie)
}

func (ms movieService) DeleteMovie(id string) error {
	return ms.movieRepo.DeleteMovie(id)
}
