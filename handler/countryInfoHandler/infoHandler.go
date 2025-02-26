package countryInfoHandler

import (
	"OBLIG_1/constants"
	"OBLIG_1/handler/linker"
	"OBLIG_1/utility"
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

func InfoHandler(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Convert ISO code to country and checking limits of cities
	isoCode, limit := utility.FormatISOandLimitOfCities(request.URL.String())
	countryName, _, countryError := utility.GetCountryNameByISO(isoCode)
	if countryError != nil {
		linker.SendErrorAsJson("Error getting country name:", countryError, w)
		return
	}

	// Create new request with new url
	restCountriesAlpha := constants.REQUEST_REST_COUNTRIES_API_ALPHA + isoCode
	countriesNowCities := constants.COUNTRUES_NOW_ALL_CITIES

	//preparing json Post-Method
	requestBody, err := json.Marshal(map[string]string{"country": countryName})
	if err != nil {
		linker.SendErrorAsJson("Error marshalling request body:", err, w)
		return
	}

	resResponce1, res1err := linker.SendGetRequest(restCountriesAlpha, *client)
	resResponce2, res2err := linker.SendPostRequest(countriesNowCities, requestBody, *client)

	if res1err != nil {
		linker.SendErrorAsJson("Error sending Get request: ", res1err, w)
		return
	}
	if res2err != nil {
		linker.SendErrorAsJson("Error sending Post request: ", res2err, w)
		return
	}

	// Decoding JSON
	countries, countryErr := buildingResponce(resResponce1, resResponce2, limit)
	if countryErr != nil {
		linker.SendErrorAsJson("Error building responce:", countryErr, w)
	}

	// Printing decoded output
	json.NewEncoder(w).Encode(countries)

}

func buildingResponce(res1 http.Response, res2 http.Response, limit int) (CountryInfoStructure, error) {
	//inserting information in structurs
	var countries []CountryInfoStructure
	decoderForCountry := json.NewDecoder(res1.Body)
	if errDecoder1 := decoderForCountry.Decode(&countries); errDecoder1 != nil {
		log.Println("Error in decoding response from country: ", errDecoder1.Error())
	}
	var cities CitiesInfoStructure
	decoderForCity := json.NewDecoder(res2.Body)
	if errDecoder2 := decoderForCity.Decode(&cities); errDecoder2 != nil {
		log.Println("Error in decoding response from city: ", errDecoder2.Error())
	}
	//puting all cities in country information
	if limit == -1 {
		countries[0].Cities = cities.Cities
	} else if limit > 0 {
		for i, city := range cities.Cities {
			if i < limit {
				countries[0].Cities = append(countries[0].Cities, city)
			}
		}
	}
	if countries[0].Cities == nil {
		return countries[0], errors.New("No cities found for this country")
	}

	return countries[0], nil
}
