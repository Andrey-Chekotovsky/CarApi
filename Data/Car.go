package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type Car struct {
	Id        int64     `json:"id"`
	Brand     string    `json:"brand" validate:"required"`
	SerialNum string    `json:"serialNum" validate"SerialNum"`
	Color     string    `json:"color" validate:"required"`
	CreatedOn time.Time `json:"-"`
	UpdatedOn time.Time `json:"-"`
}

var ErrCarNotFound = fmt.Errorf("Car not found")

type Cars []*Car

func (c *Car) FromJson(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(c)
}
func (c *Cars) toJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(c)
}

func findIndexById(id int64) (int, error) {
	low := 0
	high := len(cars) - 1
	for low <= high {
		mid := (low + high) / 2
		if cars[mid].Id == id {
			return mid, nil
		}
		if cars[mid].Id > id {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1, ErrCarNotFound
}

func GetCar(id int64) (Car, error) {
	ind, err := findIndexById(id)
	if err != nil {
		return Car{}, err
	} else {
		return *cars[ind], nil
	}
}

func GetCars() Cars {
	return cars
}

func AddCar(car *Car) {
	car.Id = getNextId()
	car.CreatedOn = time.Now()
	car.UpdatedOn = time.Now()
	cars = append(cars, car)
}

func DeleteCar(id int64) error {
	ind, err := findIndexById(id)
	if err != nil {
		return err
	}
	cars = append(cars[:ind], cars[ind+1])
	return nil
}

func Update(car Car) error {
	ind, err := findIndexById(car.Id)
	cars[ind] = &car
	if err != nil {
		return err
	} else {
		return nil
	}
}

var currentId int64 = 2

func getNextId() int64 {
	currentId++
	return currentId
}

var cars = []*Car{
	{
		Id:        1,
		Brand:     "Moskvich",
		SerialNum: "ASP-62",
		Color:     "Red",
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	},
	{
		Id:        2,
		Brand:     "Volga",
		SerialNum: "ASP-52",
		Color:     "Blue",
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	},
}
