package utils

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type AppSetting struct {
	Port string `json:"port"`
}

var (
	Setting AppSetting
)

func ReadSettings(filepath string) (AppSetting, error) {
	var AppSetting AppSetting
	doc, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("ReadFile ReadSettings filepath err is %e\n", err)
		return AppSetting, err
	}
	err = json.Unmarshal(doc, &AppSetting)
	if err != nil {
		log.Printf("ReadFile Unmarshal data err is %e\n", err)
		return AppSetting, err
	}
	return AppSetting, nil
}
