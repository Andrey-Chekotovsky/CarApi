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

func GetCar(id int64) (Car, error) {
	rows, err := db.Query("SELECT id, brand, serial_number, color FROM cars WHERE id = $1", id)
	defer rows.Close()
	if err != nil {
		return Car{}, err
	}
	car := Car{}
	rows.Next()
	rows.Scan(&car.Id, &car.Brand, &car.SerialNum, &car.Color)
	return car, nil
}

func GetCars() ([]Car, error) {
	rows, err := db.Query("SELECT id, brand, serial_number, color FROM cars;")
	defer rows.Close()
	if err != nil {
		return nil, err
	}
	c := []Car{}
	for rows.Next() {
		car := Car{}
		rows.Scan(&car.Id, &car.Brand, &car.SerialNum, &car.Color)
		c = append(c, car)
	}
	return c, nil
}

func AddCar(car *Car) error {
	command := "INSERT INTO cars(brand, serial_number, color, created_on, updated_on) VALUES ($1, $2, $3, $4, $5);"
	_, err := db.Exec(command, car.Brand, car.SerialNum, car.Color, car.CreatedOn, car.UpdatedOn)
	return err
}

func DeleteCar(id int64) error {
	command := "DELETE FROM cars WHERE id = $1;"
	_, err := db.Exec(command, id)
	if err != nil {
		return err
	} else {
		return nil
	}
}

func Update(car Car) error {
	command := "UPDATE cars SET brand=$1, serial_number=$2, color=$3, updated_on=$4 WHERE id=$5;"
	_, err := db.Exec(command, car.Brand, car.SerialNum, car.Color, car.UpdatedOn, car.Id)
	return err
}
