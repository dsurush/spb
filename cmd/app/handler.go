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

func Test(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
	writer.Header().Set(contentType, value)
	x, err := strconv.ParseFloat(request.URL.Query().Get(`x`), 64)
	if err != nil {
		log.Printf("Handler Test Parsefloat x err is %e\n", err)
		writer.WriteHeader(http.StatusBadRequest)
		err := json.NewEncoder(writer).Encode([]string{"err.x_invalid"})
		if err != nil {
			log.Println(err)
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
		}
		return
	}
	Coordinate, err := services.ReadCoordinate(filepath)
	if err != nil {
		log.Printf("Handler Test ReadData filepath err is %e\n", err)
		writer.WriteHeader(http.StatusInternalServerError)
		err := json.NewEncoder(writer).Encode([]string{"err.Server_error"})
		if err != nil {
			log.Println(err)
		}
		return
	}
	currency := Coordinate.CheckInCircile(x, y)
	AllCurrency, err := services.GetAllCurrency()
	if err != nil {
		writer.WriteHeader(http.StatusFailedDependency)
		err := json.NewEncoder(writer).Encode([]string{"err.dependency_error"})
		if err != nil {
			log.Println(err)
		}
		return
	}
	err = json.NewEncoder(writer).Encode(AllCurrency.Valute[currency])
	if err != nil {
		log.Println(err)
	}
}
