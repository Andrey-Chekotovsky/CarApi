package cars

import (
	"data"
	"net/http"
)

func (c *Cars) Delete(rw http.ResponseWriter, r *http.Request) {
	id := getCarID(r)
	err := data.DeleteCar(int64(id))
	if err != nil {
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
