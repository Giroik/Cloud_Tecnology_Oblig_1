package statusHandler

import (
	"OBLIG_1/constants"
	"OBLIG_1/handler/linker"
	"encoding/json"
	"net/http"
	"time"
)

func StatusHandler(w http.ResponseWriter, r *http.Request, startTime time.Time) {

	countriesNowAPIStatus := linker.GetAPIStatus(constants.COUNTRIES_NOW_API, constants.ENDPOINTALL)
	restCountriesAPIStatus := linker.GetAPIStatus(constants.REST_COUNTRIES_API, constants.ENDPOINTCOUNTRIES)

	status := StatusStructur{

		CountriesNowAPI:  countriesNowAPIStatus,
		RestCountriesAPI: restCountriesAPIStatus,
		Version:          "v1",
		Uptime:           int64(time.Since(startTime).Seconds()),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(status)
}
