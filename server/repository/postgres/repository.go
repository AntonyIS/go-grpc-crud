package postgres

import (
	"fmt"
	"log"
	"os"

	"github.com/AntonyIS/go-grpc-crud/server/domain"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type postgresRepository struct {
	db        *gorm.DB
	tableName string
}

func NewPostgresRepository() (domain.MovieRepository, error) {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	db, err := gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", host, port, user, dbname, password))

	if err != nil {
		log.Fatal(err)
	}

	db.LogMode(false)
	db.AutoMigrate([]domain.Movie{})
	repo := &postgresRepository{}
	repo.db = db
	repo.tableName = "Movies"

	return repo, nil

}
func (repo postgresRepository) CreateMovie(movie *domain.Movie) (*domain.Movie, error) {
	if err := repo.db.Create(&movie).Error; err != nil {
		return nil, err
	}
	return movie, nil
}

func (repo postgresRepository) GetMovie(id string) (*domain.Movie, error) {
	movie := domain.Movie{}
	if result := repo.db.Find(&movie); result.Error != nil {
		return nil, result.Error
	}
	return &movie, nil
}

func (repo postgresRepository) GetMovies() (*[]domain.Movie, error) {
	movies := []domain.Movie{}
	if result := repo.db.Find(&movies); result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return &movies, nil
}

func (repo postgresRepository) UpdateMovie(movie *domain.Movie) (*domain.Movie, error) {
	if result := repo.db.Save(movie); result.Error != nil {
		return nil, result.Error
	}
	return movie, nil

}

func (repo postgresRepository) DeleteMovie(id string) error {
	movie := domain.Movie{}

	if result := repo.db.First(&movie, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	repo.db.Delete(&movie)
	return nil
}
