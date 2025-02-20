package countryInfoHandler

import (
	"OBLIG_1/constants"
	"OBLIG_1/utility"
	"encoding/json"

	"fmt"
	"log"
	"net/http"
	"strings"
)

func InfoHandler(w http.ResponseWriter, request *http.Request) {
	// Convert ISO code to country and checking limits of cities
	isoCode, limit := utility.FormatISOandLimitOfCities(request.URL.String())
	countryName, countryError := utility.GetCountryNameByISO(isoCode)
	if countryError != nil {
		http.Error(w, countryError.Error(), http.StatusInternalServerError)
	}

	//preparing json Post-Method
	requestBody, err := json.Marshal(map[string]string{"country": countryName})
	if err != nil {
		log.Println(err)
	}

	// Create new request with new url
	restCountriesAlpha := constants.REQUEST_REST_COUNTRIES_API_ALPHA + isoCode
	countriesNowCities := constants.COUNTRUES_NOW_ALL_CITIES

	request1, errCountries := http.NewRequest(http.MethodGet, restCountriesAlpha, nil)
	if errCountries != nil {
		fmt.Errorf("Error in creating request for Countries:", errCountries.Error())
	}

	request2, errCities := http.NewRequest("POST", countriesNowCities, strings.NewReader(string(requestBody)))
	if errCities != nil {
		fmt.Errorf("Error in creating request for Countries cities:", errCities.Error())
	}

	// Setting content type -> effect depends on the service provider
	request1.Header.Add("content-type", "application/json")
	request2.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()

	// Issue request
	resResponce1, errReq1 := client.Do(request1)
	if errReq1 != nil {
		fmt.Errorf("Error in response:", errReq1.Error())
	} else if resResponce1.StatusCode != http.StatusOK {
		fmt.Errorf("Error in response:", resResponce1.Status)
	} else if resResponce1.Header.Get("content-type") != "application/json" {
		fmt.Errorf("Header structure is not application/json ", resResponce1.Status)
	}
	defer resResponce1.Body.Close()

	resResponce2, errReq2 := client.Do(request2)
	if errReq2 != nil {
		fmt.Errorf("Error in response:", errReq2.Error())
	} else if resResponce2.StatusCode != http.StatusOK {
		fmt.Errorf("Error in response:", resResponce2.Status)
	} else if resResponce2.Header.Get("content-type") != "application/json" {
		fmt.Errorf("Header structure is not application/json ", resResponce2.Status)
	}
	defer resResponce2.Body.Close()

	//Printing out status of source api
	/*fmt.Println("Status:", resResponce2.Status)
	fmt.Println("Status code:", resResponce2.StatusCode)
	fmt.Println("Content type:", resResponce2.Header.Get("content-type"))
	fmt.Println("Protocol:", resResponce2.Proto)*/

	// Decoding JSON
	countries := buildingResponce(*resResponce1, *resResponce2, limit)

	// Printing decoded output
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)

}

func buildingResponce(res1 http.Response, res2 http.Response, limit int) CountryInfoStructure {
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

	return countries[0]
}
