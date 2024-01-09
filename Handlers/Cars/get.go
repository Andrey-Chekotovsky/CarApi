package cars

import (
	"data"
	"net/http"
)

func (c *Cars) Get(rw http.ResponseWriter, r *http.Request) {
	id := getCarID(r)
	car, foundErr := data.GetCar(int64(id))
	if foundErr != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	marchErr := data.ToJSON(car, rw)
	if marchErr != nil {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
}
