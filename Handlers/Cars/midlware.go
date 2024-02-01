package cars

import (
	"context"
	"data"
	"fmt"
	"net/http"
)

func (c *Cars) MiddlewareValidateCar(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Add("Content-Type", "application/json")
		car := &data.Car{}
		fmt.Println("Start deserialisation")
		err := data.FromJSON(car, r.Body)
		fmt.Println("Start deserialisation")
		if err != nil {
			c.L.Println("[Error] Deserializing product", "error", err)

			rw.WriteHeader(http.StatusBadRequest)
			return
		}
		vErr := data.NewValidation().Validate(car)
		fmt.Println("End deserialissation")
		fmt.Println("Start Validation")
		if vErr != nil {
			c.L.Println("[ERROR] validating product", vErr)

			// return the validation messages as an array
			rw.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
		fmt.Println("End validation")
		ctx := context.WithValue(r.Context(), "Car", car)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
