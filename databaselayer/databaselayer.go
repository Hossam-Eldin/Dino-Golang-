package databaselayer

import (
	"errors"
)

const (

	MONGODB uint8 = iota
)

type DinoDBHandler interface {
	GetAvailableDynos() ([]Animal, error)
	GetDynoByNickname(string) (Animal, error)
	GetDynosByType(string) ([]Animal, error)
	AddAnimal(Animal) error
	UpdateAnimal(Animal, string) error
}

type Animal struct {
	ID         int    `bson:"-"`
	AnimalType string `bson:"animal_type"`
	Nickname   string `bson:"nickname"`
	Zone       int    `bson:"zone"`
	Age        int    `bson:"age"`
}

var DBTypeNotSupported = errors.New("The Database type provided is not supported...")

//factory function
func GetDatabaseHandler(dbtype uint8, connection string) (DinoDBHandler, error) {

	switch dbtype {

	case MONGODB:
		return NewMongodbHandler(connection)
		}
	return nil, DBTypeNotSupported
}
