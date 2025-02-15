package handler

import (
	"OBLIG_1/constants"
	"OBLIG_1/utility"
	"encoding/json"
	"net/http"
	"time"
)

func StatusHandler(w http.ResponseWriter, r *http.Request, startTime time.Time) {

	countriesNowAPIStatus := utility.GetAPIStatus(constants.COUNTRIESNOWAPI, constants.ENDPOINTALL)
	restCountriesAPIStatus := utility.GetAPIStatus(constants.RESTCOUNTRIESAPI, constants.ENDPOINTCOUNTRIES)

	status := utility.StatusStructur{

		CountriesNowAPI:  countriesNowAPIStatus,
		RestCountriesAPI: restCountriesAPIStatus,
		Version:          "v1",
		Uptime:           int64(time.Since(startTime).Seconds()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
