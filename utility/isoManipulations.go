package utility

import (
	"OBLIG_1/constants"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func FormatISOandLimitOfCities(request string) (string, int) {
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
func FormatISOandPopulationYears(request string) (string, int, int) {

	parsedURL, err := url.Parse(request)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}
	path := parsedURL.Path
	splitedPath := strings.Split(path, "/")
	iso := splitedPath[len(splitedPath)-1]
	limitString := parsedURL.Query().Get("limit")
	years := strings.Split(limitString, "-")
	if len(years) == 1 {
		startYear, syerr := strconv.Atoi(years[0])
		if syerr != nil {
			startYear = 0
		}
		return iso, startYear, 0
	} else if len(years) == 2 {
		startYear, syerr := strconv.Atoi(years[0])
		if syerr != nil {
			startYear = 0
		}
		endYear, eyerr := strconv.Atoi(years[1])
		if eyerr != nil {
			endYear = 0
		}
		return iso, startYear, endYear
	}

	fmt.Println(limitString)
	return iso, 0, 0
}

func GetCountryNameByISO(userIso string) (string, error) {
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

// reserv version of getCountryNameByISO if county was not found in default version
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
