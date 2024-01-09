package cars

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Cars struct {
	L *log.Logger
}

func getCarID(r *http.Request) int64 {
	// parse the product id from the url
	vars := mux.Vars(r)

	// convert the id into an integer and return
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		// should never happen
		panic(err)
	}

	return id
}
