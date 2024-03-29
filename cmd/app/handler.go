package app

import (
	"encoding/json"
	"log"
	"net/http"
	"spb/pkg/services"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

const contentType = "Content-Type"
const value = "application/json; charset=utf-8"
const filepath = `../data.json`

var coordinate services.CircleCoordinateData

func InradiusHandler(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {

	writer.Header().Set(contentType, value)
	x, err := strconv.ParseFloat(request.URL.Query().Get(`x`), 64)

	if err != nil {
		log.Printf("Handler Test Parsefloat x err is %e\n", err)
		writer.WriteHeader(http.StatusBadRequest)

		err := json.NewEncoder(writer).Encode([]string{"err.x_invalid"})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	y, err := strconv.ParseFloat(request.URL.Query().Get(`y`), 64)
	if err != nil {
		log.Printf("Handler Test Parsefloat y err is %e\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.y_invalid"})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	allCurrency, err := services.GetAllCurrency()
	if err != nil {
		writer.WriteHeader(http.StatusFailedDependency)
		err := json.NewEncoder(writer).Encode([]string{"err.dependency_error"})
		if err != nil {
			log.Println(err)
			return
		}
		return
	}

	currency := coordinate.CheckInCircile(x, y)
	err = json.NewEncoder(writer).Encode(allCurrency.Valute[currency])
	if err != nil {
		log.Println(err)
		return
	}
}
