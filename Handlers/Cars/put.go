package cars

import (
	"data"
	"net/http"
)

func (c *Cars) Put(rw http.ResponseWriter, r *http.Request) {
	car := r.Context().Value("Car").(*data.Car)
	err := data.Update(*car)
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
