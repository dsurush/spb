package services

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

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
	url := `https://www.cbr-xml-daily.ru/daily_json.js`
	method := `GET`

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Printf("Can't build request err is %e\n", err)
		return data, err
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Can't send Request err is %e\n", err)
		return data, err
	}

	if resp.StatusCode != http.StatusOK {
		denidedError := errors.New("server is denided")
		return data, denidedError

	}

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
}
