package countryInfoHandler

import (
	"OBLIG_1/constants"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func InfoHandler(w http.ResponseWriter, request *http.Request) {
	// Convert ISO code to country and checking limits of cities
	isoCode, limit := formatISOandLimit(request.URL.String())
	countryName, countryError := getCountryNameByISO(isoCode)
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

func getCountryNameByISO(userIso string) (string, error) {
	userCountry := "notFound" //init string to return

	//Creating url for get-request
	isoURL := constants.REQUEST_ISO_CODES

	//creating request to get all ISO
	isoRequest, reqIsoErr := http.NewRequest(http.MethodGet, isoURL, nil)
	if reqIsoErr != nil {
		fmt.Errorf("Error in creating request for iso codes:", reqIsoErr.Error())
	}

	// Creating independent client to get all iso codes
	isoClient := &http.Client{}
	defer isoClient.CloseIdleConnections()

	//Getting Responce from server
	isoResponse, isoErr := isoClient.Do(isoRequest)
	if isoErr != nil {
		fmt.Errorf("Error in response to iso:", isoErr.Error())
	} else if isoResponse.StatusCode != http.StatusOK {
		fmt.Errorf("Error in response to iso:", isoResponse.Status)
	} else if isoResponse.Header.Get("content-type") != "application/json" {
		fmt.Errorf("Header structure is not application/json ", isoResponse.Status)
	}
	defer isoResponse.Body.Close()

	//Creating  structure for ISO
	var isoCountryInfo CountriesNowISOCountries
	isoDecoder := json.NewDecoder(isoResponse.Body)
	if errDecoder := isoDecoder.Decode(&isoCountryInfo); errDecoder != nil {
		fmt.Errorf("Error in decoding response from iso codes:", errDecoder.Error())
	} else {
		fmt.Println("All ", len(isoCountryInfo.Data), " iso codes are successful found")
	}
	//searching for right country
	for _, country := range isoCountryInfo.Data {
		if strings.EqualFold(country.Iso2, userIso) || strings.EqualFold(country.Iso3, userIso) {
			userCountry = country.Name
		}
	}
	//If not found run reserve API to confirm not existing
	if userCountry == constants.NOT_FOUND {
		userCountry, _ = getReserveCountryNameByISO(userIso) //runs reserve api and tries to find ISO there
		if userCountry == constants.NOT_FOUND {
			return userCountry, errors.New("User country not found") //returns error
		}
	}
	return userCountry, nil
}

// reserv version of getCountryNameByISO
func getReserveCountryNameByISO(iso string) (string, error) {
	userCountry := constants.NOT_FOUND                      //init return string
	reservIsoURL := constants.RESERV_REQUEST_ISO_CODE + iso //preapering url

	//Creating request
	reservIsoRequest, resReqIsoErr := http.NewRequest(http.MethodGet, reservIsoURL, nil)
	if resReqIsoErr != nil {
		fmt.Errorf("Error in creating request for iso codes:", resReqIsoErr.Error())
	}

	//creating reserv Client
	reservIsoClient := &http.Client{}
	defer reservIsoClient.CloseIdleConnections()

	//Geting responce
	reservIsoResponce, resIsoErr := reservIsoClient.Do(reservIsoRequest)
	if resIsoErr != nil {
		fmt.Errorf("Error in response to reserve iso:", resIsoErr.Error())
	} else if reservIsoResponce.StatusCode != http.StatusOK {
		fmt.Errorf("Error in response to reserve iso:", reservIsoResponce.Status)
	} else if reservIsoResponce.Header.Get("content-type") != "application/json" {
		fmt.Errorf("Header structure is not application/json ", reservIsoResponce.Status)
	}
	defer reservIsoResponce.Body.Close()

	//creating structure for other API responce
	var reserveIsoCheck []CountryIsoCheck
	reserveIsoDecoder := json.NewDecoder(reservIsoResponce.Body)
	if errDecoder := reserveIsoDecoder.Decode(&reserveIsoCheck); errDecoder != nil {
		fmt.Errorf("Error in decoding response from reserve iso:", errDecoder.Error())
	}
	//searching for right country
	for _, country := range reserveIsoCheck {
		if strings.EqualFold(country.Cca2, iso) || strings.EqualFold(country.Cca3, iso) {
			userCountry = country.Name.Common
		}
	}
	//if not existing return error
	if userCountry == constants.NOT_FOUND {
		return userCountry, errors.New("User country not found")
	}
	return userCountry, nil
}

func formatISOandLimit(request string) (string, int) {
	parsedURL, err := url.Parse(request)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}
	path := parsedURL.Path
	splitedPath := strings.Split(path, "/")
	iso := splitedPath[len(splitedPath)-2]
	limitString := parsedURL.Query().Get("limit")
	limit, err := strconv.Atoi(limitString)
	if err != nil {
		fmt.Println("Error parsing limit:", err)
		limit = -1
	}

	return iso, limit
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
