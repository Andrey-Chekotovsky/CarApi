package cars

import (
	"data"
	"net/http"
)

func (c *Cars) GetAll(rw http.ResponseWriter, r *http.Request) {
	cars, err1 := data.GetCars()
	if err1 != nil {
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	err2 := data.ToJSON(cars, rw)
	if err2 != nil {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
}
