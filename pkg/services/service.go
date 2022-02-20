package services

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type CircleCoordinateData struct {
	X       float64 `json:"x"`
	Y       float64 `json:"y"`
	Diametr float64 `json:"diametr"`
}

const usd = "USD"
const eur = "EUR"

func ReadCoordinate(filepath string) (CircleCoordinateData, error) {

	var сoordinate CircleCoordinateData
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

//По теореме пифагора находим расстояние между точками, но при этом не берем корень.
func (receiver *CircleCoordinateData) сountDistance(x, y float64) float64 {
	distance := (receiver.X-x)*(receiver.X-x) + (receiver.Y-y)*(receiver.Y-y)
	return distance
}

func (receiver *CircleCoordinateData) CheckInCircile(x, y float64) string {

	distance := receiver.сountDistance(x, y)
	radius := receiver.Diametr / 2
	//Так как эти значения имеют тип float
	//было бы правильно сравнивать разницу абсолютного значения с epsilon
	//radius * radius - distance >= epsilon
	if radius*radius >= distance {
		return usd
	}
	return eur
}
