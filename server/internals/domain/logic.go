package domain

import (
	"errors"

	"github.com/teris-io/shortid"
)

var (
	ErrItemNotFound        = errors.New("item not found")
	ErrInternalServerError = errors.New("internal server error")
)

type carService struct {
	carRepo CarRepository
}

func NewCarService(carRepo CarRepository) CarService {
	return &carService{
		carRepo,
	}
}

func (cs *carService) CreateCar(car *Car) (*Car, error) {
	car.ID = shortid.MustGenerate()
	return cs.carRepo.CreateCar(car)
}

func (cs *carService) GetCar(id string) (*Car, error) {
	return cs.carRepo.GetCar(id)
}

func (cs *carService) GetCars() (*[]Car, error) {
	return cs.carRepo.GetCars()
}

func (cs *carService) UpdateCar(car *Car) (*Car, error) {
	return cs.carRepo.UpdateCar(car)
}

func (cs *carService) DeleteCar(id string) error {
	return cs.carRepo.DeleteCar(id)
}
