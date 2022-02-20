package app

import (
	"fmt"
	"log"
	"net/http"
	"spb/pkg/services"
	"spb/settings/utils"

	"github.com/julienschmidt/httprouter"
)

func RouterInit() {
	router := httprouter.New()

	setting, err := utils.ReadSettings(`../settings/settings.json`)
	if err != nil {
		log.Fatalf("Check path envirement readsettings err is %e", err)
	}

	coordinate, err = services.ReadCoordinate(filepath)
	if err != nil {
		log.Fatalf("Check path bussines readsettings err is %e", err)
	}

	router.GET("/inradius", InradiusHandler)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", setting.Port), router))
}
