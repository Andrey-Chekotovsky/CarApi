package cars

import (
	"data"
	"net/http"
)

func (c *Cars) Put(rw http.ResponseWriter, r *http.Request) {
	car := data.Car{}
	err := car.FromJson(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	data.Update(car)
	rw.WriteHeader(http.StatusOK)
}
