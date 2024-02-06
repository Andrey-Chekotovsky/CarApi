package cars

import (
	"data"
	"fmt"
	"net/http"
	"time"
)

func (c *Cars) Put(rw http.ResponseWriter, r *http.Request) {
	o := r.Context().Value("Car")
	car := o.(*data.Car)
	car.UpdatedOn = time.Now()
	err := data.Update(*car)
	if err != nil {
		fmt.Println("[ERROR] Can't add")
		rw.WriteHeader(http.StatusNotFound)
		return
	}
	rw.WriteHeader(http.StatusOK)
}
