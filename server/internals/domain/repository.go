package domain

type CarRepository interface {
	CreateCar(car *Car) (*Car, error)
	GetCar(id string) (*Car, error)
	GetCars() (*[]Car, error)
	UpdateCar(car *Car) (*Car, error)
	DeleteCar(id string) error
}
