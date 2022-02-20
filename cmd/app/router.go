package app

import (
	"fmt"
	"log"
	"net/http"
	"spb/settings"

	"github.com/julienschmidt/httprouter"
)

func RouterInit() {
	router := httprouter.New()
	setting, err := settings.ReadSettings(`../settings/settings.json`)
	if err != nil {
		log.Fatalf("Check path readsettings")
	}
	router.GET("/test", Test)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", setting.Port), router))
}
