package cars

import (
	"data"
	"net/http"
)

func (c *Cars) GetAll(rw http.ResponseWriter, r *http.Request) {
	err := data.ToJSON(data.GetCars(), rw)
	if err != nil {
		rw.WriteHeader(http.StatusBadRequest)
	} else {
		rw.WriteHeader(http.StatusOK)
	}
}
