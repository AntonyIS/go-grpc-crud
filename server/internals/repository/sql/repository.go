package sql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/AntonyIS/go-grpc-crud/server/internals/domain"
)

type sqlRepository struct {
	db        *sql.DB
	tablename string
}

func dbClient() *sql.DB {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "gofoods"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	// Ping db
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	return db
}
func NewRepository(tablename string) (domain.CarRepository, error) {
	repo := &sqlRepository{}
	repo.db = dbClient()
	return repo, nil
}

func (sql sqlRepository) CreateCar(car *domain.Car) (*domain.Car, error) {
	operation := fmt.Sprintf("INSERT INTO %v", sql.tablename)
	values := fmt.Sprintf("VALUES (%v %v %v) ", car.ID, car.Name, car.Description)
	insert, err := sql.db.Query(operation + sql.tablename + values)

	if err != nil {
		panic(err.Error())
	}
	defer insert.Close()
	return car, nil
}

func (sql sqlRepository) GetCars() (*[]domain.Car, error) {
	query, err := sql.db.Query(fmt.Sprintf("SELECT * FROM %v", sql.tablename))

	if err != nil {
		log.Fatalf("Error : %v", err)
	}
	car := domain.Car{}
	cars := []domain.Car{}

	for query.Next() {
		var id, name, description string
		err = query.Scan(&id, &name, &description)

		if err != nil {
			panic(err.Error())
		}
		car.ID = id
		car.Name = name
		car.Description = description

		cars = append(cars, car)
	}
	defer query.Close()
	defer sql.db.Close()
	return &cars, nil

}

func (sql sqlRepository) GetCar(id string) (*domain.Car, error) {
	car := domain.Car{}
	query, err := sql.db.Query("SELECT * FROM Employee WHERE id=?", id)
	if err != nil {
		panic(err.Error())
	}

	if query.Next() {
		err = query.Scan(&car.ID, &car.Name, &car.Description)
		if err != nil {
			panic(err.Error())
		}
	}
	defer query.Close()
	defer sql.db.Close()
	return &car, nil

}

func (sql sqlRepository) UpdateCar(car *domain.Car) (*domain.Car, error) {
	prep, err := sql.db.Prepare(fmt.Sprintf("UPDATE %v SET name=?, email=? WHERE user_id=?", sql.tablename))

	if err != nil {
		panic(err.Error())
	}
	prep.Exec(car.Name, car.Description)

	defer sql.db.Close()
	return car, nil
}

func (sql sqlRepository) DeleteCar(id string) error {
	query, err := sql.db.Prepare(fmt.Sprintf("DELETE FROM %v WHERE id=?", sql.tablename))
	if err != nil {
		panic(err.Error())
	}
	query.Exec(id)
	defer sql.db.Close()
	return nil
}
