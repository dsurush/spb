package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type CoordinateData struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Diametr float64 `json:"diametr"`
}

const USD = "USD"
const EUR = "EUR"

func ReadCoordinate(filepath string) (CoordinateData, error) {

	var сoordinate CoordinateData
	doc, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("ReadFile ReadSettings filepath err is %e\n", err)
		return сoordinate, err
	}

	err = json.Unmarshal(doc, &сoordinate)
	if err != nil {
		log.Printf("ReadFile Unmarshal data err is %e\n", err)
		return сoordinate, err
	}

	return сoordinate, nil
}

func (receiver *CoordinateData) сountDistance(x, y float64) float64 {
	distance := (receiver.X-x)*(receiver.X-x) + (receiver.Y-y)*(receiver.Y-y)
	return distance
}

func (receiver *CoordinateData) CheckInCircile(x, y float64) string {

	distance := receiver.сountDistance(x, y)
	radius := receiver.Diametr / 2

	if radius*radius >= distance {
		return USD
	}
	return EUR
}
