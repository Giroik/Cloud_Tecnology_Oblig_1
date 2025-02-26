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
	if limitString != "" {
		limit, err := strconv.Atoi(limitString)
		if err != nil {
			fmt.Println("Error parsing limit:", err)
			return iso, constants.RESET_LIMIT
		}
		return iso, limit
	}

	return iso, constants.RESET_LIMIT
}
func FormatISOandPopulationYears(request string) (string, int, int, error) {

	parsedURL, err := url.Parse(request)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
	}
	path := parsedURL.Path
	splitedPath := strings.Split(path, "/")
	iso := splitedPath[len(splitedPath)-1]
	limitString := parsedURL.Query().Get("limit")
	if limitString != "" {
		years := strings.Split(limitString, "-")
		if len(years) == 1 {
			startYear, syerr := strconv.Atoi(years[0])
			if syerr != nil {
				startYear = 0
			}
			return iso, startYear, -1, nil
		} else if len(years) == 2 {
			startYear, syerr := strconv.Atoi(years[0])
			if syerr != nil {
				startYear = 0
			}
			endYear, eyerr := strconv.Atoi(years[1])
			if eyerr != nil {
				endYear = 0
			}
			return iso, startYear, endYear, nil
		} else if len(years) > 2 {
			return iso, -1, -1, errors.New("Too many requests in ?limit=")
		}
	}
	return iso, 0, 0, nil
}

func GetReserveCountryNameByISO(userIso string) (string, error) {
	userCountry := constants.NOT_FOUND //init string to return

	//Creating url for get-request
	reserveIsoURL := constants.RESERVE_REQUEST_ISO_CODES

	//creating request to get all ISO
	isoRequest, reqIsoErr := http.NewRequest(http.MethodGet, reserveIsoURL, nil)
	if reqIsoErr != nil {
		return "", errors.New("Error creating reserve request for " + reserveIsoURL + ": " + reqIsoErr.Error())
	}
	isoRequest.Header.Add("Content-Type", "application/json")

	// Creating independent client to get all iso codes
	isoClient := &http.Client{}
	defer isoClient.CloseIdleConnections()

	//Getting Responce from server
	isoResponse, isoErr := isoClient.Do(isoRequest)
	if isoErr != nil {
		fmt.Println("Error in response to reserve iso:", isoErr.Error())
	} else if isoResponse.StatusCode != http.StatusOK {
		fmt.Println("Error in response to reserve iso:", isoResponse.Status)
	} else if isoResponse.Header.Get("content-type") != "application/json" {
		fmt.Println("Header structure is not application/json ", isoResponse.Status)
	}
	defer isoResponse.Body.Close()

	//Creating  structure for ISO
	var isoCountryInfo CountriesNowISOCountries
	isoDecoder := json.NewDecoder(isoResponse.Body)
	if errDecoder := isoDecoder.Decode(&isoCountryInfo); errDecoder != nil {
		fmt.Println("Error in decoding response from reserve iso codes:", errDecoder.Error())
	} else {
		fmt.Println("All ", len(isoCountryInfo.Data), "reserve iso codes are successful found")
	}
	//searching for right country
	for _, country := range isoCountryInfo.Data {
		if strings.EqualFold(country.Iso2, userIso) || strings.EqualFold(country.Iso3, userIso) {
			userCountry = country.Name
		}
	}
	//If not found run reserve API to confirm not existing

	return userCountry, nil
}

// reserv version of getCountryNameByISO if county was not found in default version
func GetCountryNameByISO(iso string) (string, string, error) {
	userCountry := constants.NOT_FOUND //init return string
	userCountryOfficial := constants.NOT_FOUND
	isoURL := constants.REQUEST_ISO_CODE + iso //preapering url

	//Creating request
	IsoRequest, ReqIsoErr := http.NewRequest(http.MethodGet, isoURL, nil)
	if ReqIsoErr != nil {
		fmt.Println("Error in creating request for iso codes:", ReqIsoErr.Error())
	}

	//creating reserv Client
	isoClient := &http.Client{}
	defer isoClient.CloseIdleConnections()

	//Geting responce
	IsoResponce, IsoErr := isoClient.Do(IsoRequest)
	if IsoErr != nil {
		fmt.Println("Error in response iso:", IsoErr.Error())
		fmt.Println(IsoResponce.Body)
	} else if IsoResponce.StatusCode != http.StatusOK {
		fmt.Println("Error in response iso:", IsoResponce.Status)
	} else if IsoResponce.Header.Get("content-type") != "application/json" {
		fmt.Println("Header structure is not application/json ", IsoResponce.Status)
	}
	defer IsoResponce.Body.Close()

	//creating structure for other API responce
	var IsoCheck []CountryIsoCheck
	IsoDecoder := json.NewDecoder(IsoResponce.Body)
	if errDecoder := IsoDecoder.Decode(&IsoCheck); errDecoder != nil {
		fmt.Println("Error in decoding response from iso:", errDecoder.Error())
	}
	//searching for right country
	for _, country := range IsoCheck {
		if strings.EqualFold(country.Cca2, iso) || strings.EqualFold(country.Cca3, iso) {
			userCountry = country.Name.Common
			userCountryOfficial = country.Name.Official
		}
	}
	//if not existing return error
	if userCountry == constants.NOT_FOUND {
		//runs reserve api and tries to find ISO there
		countryName, myerror := GetReserveCountryNameByISO(iso)
		if myerror != nil {
			return "", "", myerror
		}
		userCountry = countryName
		if countryName == constants.NOT_FOUND {
			return "", "", errors.New("Contry with code iso: " + iso + " not found")
		}
	}
	return userCountry, userCountryOfficial, nil
}
