package cars

import (
	"data"
	"fmt"
	"net/http"
	"time"
)

func (c *Cars) AddCar(rw http.ResponseWriter, r *http.Request) {
	o := r.Context().Value("Car")
	car := o.(*data.Car)
	car.CreatedOn = time.Now()
	car.UpdatedOn = time.Now()
	err := data.AddCar(car)
	if err != nil {
		fmt.Println("CAn't add")
		panic(err)
	}
	rw.WriteHeader(http.StatusCreated)
}
