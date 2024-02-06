package data

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"time"

	_ "github.com/lib/pq"
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

var db *sql.DB

func ConnectToDb(connstring string) error {
	var err error
	db, err = sql.Open("postgres", connstring)
	return err
}

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
