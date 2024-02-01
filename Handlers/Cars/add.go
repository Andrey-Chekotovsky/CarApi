package cars

import (
	"data"
	"net/http"
)

func (c *Cars) AddCar(rw http.ResponseWriter, r *http.Request) {
	car := r.Context().Value("Car")
	data.AddCar(car.(*data.Car))
	rw.WriteHeader(http.StatusCreated)
}
