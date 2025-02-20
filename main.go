package main

import (
	"OBLIG_1/handler"
	"OBLIG_1/handler/countryInfoHandler"
	"OBLIG_1/handler/populationHandler"
	"OBLIG_1/handler/statusHandler"
	"log"
	"net/http"
	"os"
	"time"
)

var startTime time.Time

func main() {
	startTime = time.Now()

	PORT := "8080"

	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}

	router := http.NewServeMux()

	router.HandleFunc("/countryinfo/v1/", handler.FrontPageHandler)
	router.HandleFunc("/countryinfo/v1/info/{country_code}/", countryInfoHandler.InfoHandler)
	router.HandleFunc("/countryinfo/v1/population/{country_code}", populationHandler.PopulationHandler)
	router.HandleFunc("/countryinfo/v1/status/", func(w http.ResponseWriter, r *http.Request) {
		statusHandler.StatusHandler(w, r, startTime)
	})
	//router.HandleFunc("/bober/{param1...}/", handler.DiagHandler)

	log.Println("Starting server on port " + PORT)
	log.Fatal(http.ListenAndServe(":"+PORT, router))

}
