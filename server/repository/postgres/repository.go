package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/AntonyIS/go-grpc-crud/server/domain"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type postgresRepository struct {
	db        *sql.DB
	tableName string
}

func NewPostgresRepository() (domain.MovieRepository, error) {
	repo := &postgresRepository{}
	err := godotenv.Load(".env")

	if err != nil {
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}
	host := os.Getenv("HOST")
	port := 5432
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", conn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to the database")
	repo.db = db
	return repo, nil

}
func (repo postgresRepository) CreateMovie(movie *domain.Movie) (*domain.Movie, error) {
	defer repo.db.Close()

	insert := fmt.Sprintf("INSERT INTO %s values ('%s','%s','%s','%s','%s')", repo.tableName, movie.ID, movie.Name, movie.Description, movie.ReleaseDate, movie.Image)
	_, err := repo.db.Exec(insert)

	if err != nil {
		return nil, err
	}
	return movie, nil
}

func (repo postgresRepository) GetMovie(id string) (*domain.Movie, error) {
	defer repo.db.Close()
	movie := domain.Movie{}
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", repo.tableName)
	row := repo.db.QueryRow(query, id)

	err := row.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.ReleaseDate, &movie.Image)

	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (repo postgresRepository) GetMovies() (*[]domain.Movie, error) {

	query := fmt.Sprint("SELECT * FTOM %s", repo.tableName)

	rows, err := repo.db.Query(query)

	if err != nil {
		return nil, err
	}
	repo.db.Close()
	movies := []domain.Movie{}

	for rows.Next() {
		movie := domain.Movie{}

		err := rows.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.ReleaseDate, &movie.Image)
		if err != nil {
			return nil, err
		}
		movies = append(movies, movie)
	}

	return &movies, nil
}

func (repo postgresRepository) UpdateMovie(movie *domain.Movie) (*domain.Movie, error) {
	defer repo.db.Close()
	update := fmt.Sprintf(" `UPDATE %s SET name=$2 description=$3 release_date=$4 image=$5 WHERE id=$1`", repo.tableName)
	_, err := repo.db.Exec(update, movie.ID, movie.Name, movie.Description, movie.Image)

	if err != nil {
		return nil, err
	}
	return movie, nil

}

func (repo postgresRepository) DeleteMovie(id string) error {
	defer repo.db.Close()
	del := fmt.Sprintf("DELETE FROM %s WHERE id=$1")

	_, err := repo.db.Exec(del, id)
	if err != nil {
		return err
	}
	return nil
}
