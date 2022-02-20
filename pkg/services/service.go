package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

type CoordinateData struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Diametr float64 `json:"diametr"`
}

func ReadCoordinate(filepath string) (CoordinateData, error) {
	var Coordinate CoordinateData
	doc, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("ReadFile ReadSettings filepath err is %e\n", err)
		return Coordinate, err
	}
	err = json.Unmarshal(doc, &Coordinate)
	if err != nil {
		log.Printf("ReadFile Unmarshal data err is %e\n", err)
		return Coordinate, err
	}
	return Coordinate, nil
}

func (receiver *CoordinateData) CountDistance(x, y float64) float64 {
	distance := (receiver.X-x)*(receiver.X-x) + (receiver.Y-y)*(receiver.Y-y)
	return distance
}

func (receiver *CoordinateData) CheckInCircile(x, y float64) string {
	const USD = "USD"
	const EUR = "EUR"

	distance := receiver.CountDistance(x, y)
	radius := receiver.Diametr / 2

	if radius*radius >= distance {
		return USD
	}
	return EUR
}

type AllCurrency struct {
	Date         string              `json:"Date"`
	PreviousDate string              `json:"PreviousDate"`
	PreviousURL  string              `json:"PreviousURL"`
	Timestamp    string              `json:"Timestamp"`
	Valute       map[string]Currency `json:"Valute"`
}

type Currency struct {
	ID       string  `json:"ID"`
	NumCode  string  `json:"NumCode"`
	CharCode string  `json:"CharCode"`
	Nominal  int64   `json:"Nominal"`
	Name     string  `json:"Name"`
	Value    float64 `json:"Value"`
	Previous float64 `json:"Previous"`
}

func GetAllCurrency() (AllCurrency, error) {
	client := &http.Client{}
	data := AllCurrency{}
	URL := `https://www.cbr-xml-daily.ru/daily_json.js`
	Method := `GET`
	req, err := http.NewRequest(Method, URL, nil)
	if err != nil {
		log.Printf("Can't build request err is %e\n", err)
		return data, err
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Can't send Request err is %e\n", err)
		return data, err
	}
	if resp.StatusCode == http.StatusOK {
		allData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Can't read resp.Body err is %e\n", err)
			return data, err
		}
		err = json.Unmarshal(allData, &data)
		if err != nil {
			log.Printf("json invalid err = %e\n", err)
			return data, err
		}
		return data, nil
	} else {
		denidedError := errors.New("server is denided")
		return data, denidedError
	}
}
