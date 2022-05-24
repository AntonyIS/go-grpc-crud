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
func (repo postgresRepository) CreateMovie(movie domain.Movie) (*domain.Movie, error) {
	defer repo.db.Close()
	sqlStatement := `INSERT INTO movies (id,name, description,release_date, image) VALUES ($1, $2, $3, $4, $5)`
	err := repo.db.QueryRow(sqlStatement, movie.ID, movie.Name, movie.Description, movie.ReleaseDate)

	if err != nil {
		return nil, err
	}
	return &movie, nil
}

func (repo postgresRepository) GetMovie(id string) (*domain.Movie, error) {
	defer repo.db.Close()
	movie := *&domain.Movie{}
	sqlStatement := `SELECT * FROM movies WHERE userid=$1`
	row := repo.db.QueryRow(sqlStatement, id)
	err := row.Scan(&movie.ID, &movie.Name, &movie.Description, &movie.Image)

	switch err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return &movie, nil
	case nil:
		return &movie, nil
	default:
		log.Fatalf("Unable to scan the row. %v", err)
	}
	return &movie, nil
}
