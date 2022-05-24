package domain

type MovieService interface {
	CreateMovie(movie *Movie) (*Movie, error)
	GetMovie(id string) (*Movie, error)
	GetMovies() (*[]Movie, error)
	UpdateMovie(movie *Movie) (*Movie, error)
	DeleteMovie(id string) error
}
