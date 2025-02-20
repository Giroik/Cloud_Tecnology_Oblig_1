package countryInfoHandler

import (
	"OBLIG_1/constants"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func InfoHandler(w http.ResponseWriter, request *http.Request) {
	// Convert ISO code to country
	isoCode := formatISO(request.URL.Path, constants.CONTRY_INFORMATION)
	countryName, countryError := checkCountryISO(isoCode)
	if countryError != nil {
		http.Error(w, countryError.Error(), http.StatusInternalServerError)
	}

	var jsonStr = []byte(`{"country":"` + countryName + `"}`)
	requestBody, err := json.Marshal(map[string]string{"country": countryName})
	if err != nil {
		log.Println(err)
	}
	println(string(jsonStr))

	// Create new request with new url
	restCountrysAlpha := constants.REQUEST_REST_COUNTRIES_API_ALPHA + isoCode + "/"

	countrysNowCities := constants.COUNTRUES_NOW_ALL_CITIES

	request1, errCountries := http.NewRequest(http.MethodGet, restCountrysAlpha, nil)
	if errCountries != nil {
		fmt.Errorf("Error in creating request for Countries:", errCountries.Error())
	}

	request2, errCities := http.NewRequest("POST", countrysNowCities, strings.NewReader(string(requestBody)))
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
	resRequest1, errReq1 := client.Do(request1)
	if errReq1 != nil {
		fmt.Errorf("Error in response:", errReq1.Error())
	} else if resRequest1.StatusCode != http.StatusOK {
		fmt.Errorf("Error in response:", resRequest1.Status)
	} else if resRequest1.Header.Get("content-type") != "application/json" {
		fmt.Errorf("Header structure is not application/json ", resRequest1.Status)
	}
	defer resRequest1.Body.Close()

	resRequest2, errReq2 := client.Do(request2)
	if errReq2 != nil {
		fmt.Errorf("Error in response:", errReq2.Error())
	} else if resRequest2.StatusCode != http.StatusOK {
		fmt.Errorf("Error in response:", resRequest2.Status)
	} else if resRequest2.Header.Get("content-type") != "application/json" {
		fmt.Errorf("Header structure is not application/json ", resRequest2.Status)
	}
	defer resRequest2.Body.Close()

	//Printing out status
	fmt.Println("Status:", resRequest2.Status)
	fmt.Println("Status code:", resRequest2.StatusCode)
	fmt.Println("Content type:", resRequest2.Header.Get("content-type"))
	fmt.Println("Protocol:", resRequest2.Proto)

	/*output, err := io.ReadAll(resRequest2.Body)
	if err != nil {
		fmt.Errorf("Error when reading response: ", err.Error())
	}
	fmt.Println(string(output))*/

	// Decoding JSON
	var countries []CountryInfoStructure
	decoderForCountry := json.NewDecoder(resRequest1.Body)
	if errDecoder1 := decoderForCountry.Decode(&countries); errDecoder1 != nil {
		log.Println("Error in decoding response from country: ", errDecoder1.Error())
	}
	var cities CitiesInfoStructure
	decoderForCity := json.NewDecoder(resRequest2.Body)
	if errDecoder2 := decoderForCity.Decode(&cities); errDecoder2 != nil {
		log.Println("Error in decoding response from city: ", errDecoder2.Error())
	}
	//puting all cities in country information
	countries[0].Cities = cities.Cities

	// Printing decoded output
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(countries)

}

func checkCountryISO(userIso string) (string, error) {
	fmt.Println(userIso)
	userCountry := ""

	isoURL := constants.REQUEST_ISO_CODES
	isoRequest, isoErr := http.NewRequest(http.MethodGet, isoURL, nil)
	isoClient := &http.Client{}
	defer isoClient.CloseIdleConnections()

	if isoErr != nil {
		fmt.Errorf("Error in creating request for iso codes:", isoErr.Error())
	}

	isoResponse, isoErr := isoClient.Do(isoRequest)
	if isoErr != nil {
		fmt.Errorf("Error in response to iso:", isoErr.Error())
	} else if isoResponse.StatusCode != http.StatusOK {
		fmt.Errorf("Error in response to iso:", isoResponse.Status)
	} else if isoResponse.Header.Get("content-type") != "application/json" {
		fmt.Errorf("Header structure is not application/json ", isoResponse.Status)
	}
	defer isoResponse.Body.Close()

	var isoCountryInfo CountriesNowISOCountries
	isoDecoder := json.NewDecoder(isoResponse.Body)
	if errDecoder := isoDecoder.Decode(&isoCountryInfo); errDecoder != nil {
		fmt.Errorf("Error in decoding response from iso codes:", errDecoder.Error())
	} else {
		count := 0
		for range isoCountryInfo.Data {
			count++
		}
		fmt.Println("All ", count, " iso codes are successful found")
	}
	//searching for right country
	for _, country := range isoCountryInfo.Data {
		if strings.EqualFold(country.Iso2, userIso) || strings.EqualFold(country.Iso3, userIso) {
			userCountry = country.Name
		}
	}
	if userCountry == "" {
		return userCountry, errors.New("User country not found")
	}
	return userCountry, nil
}

func formatISO(request string, toDelete string) string {
	splited := strings.Split(request, "/")
	iso := splited[len(splited)-2]

	return iso
}

func checkLimit(url string) int {
	return 0
}
