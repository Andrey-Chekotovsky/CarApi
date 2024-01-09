package cars

import (
	"data"
	"net/http"
)

func (c *Cars) AddCar(rw http.ResponseWriter, r *http.Request) {
	car := data.Car{}
	err := car.FromJson(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	data.AddCar(&car)
	rw.WriteHeader(http.StatusCreated)
}
