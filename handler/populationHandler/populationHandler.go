package populationHandler

import (
	"OBLIG_1/constants"
	"OBLIG_1/handler/linker"
	"OBLIG_1/utility"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//creating url for postRequest
	postURL := constants.REQUEST_FILTERD_POPULATION

	//Preparing client for tasks
	client := &http.Client{}
	defer client.CloseIdleConnections()

	//extracting information from users URL
	isoCode, startY, endY, yearErr := utility.FormatISOandPopulationYears(r.URL.String())
	if yearErr != nil {
		linker.SendErrorAsJson("Error parsing ISO and population years:", yearErr, w)
		fmt.Printf("Error parsing ISO and population years: %s /n Format have to be info/{iso_code}?limit=year - year", yearErr.Error())
		return
	}

	//extracting name to country throw ISO code
	countryName, officialName, countryError := utility.GetCountryNameByISO(isoCode)
	if countryError != nil {
		linker.SendErrorAsJson("Error getting country name:", countryError, w)
		fmt.Printf("Error getting country name: %s", countryError.Error())
		return
	}

	//preapering request
	requestBody, err1 := json.Marshal(map[string]string{"country": countryName})
	if err1 != nil {
		log.Println(err1)
	}
	//reserve request if api usin official nameof country
	reserveRequestBody, err2 := json.Marshal(map[string]string{"country": officialName})
	if err2 != nil {
		log.Println(err2)
	}

	// getting responce from API first with common name and if don't work, with official name
	populationResponce, popError := linker.SendPostRequest(postURL, requestBody, *client)
	if popError != nil {
		fmt.Println(w, "Error posting population response: ", popError.Error())
		ReservePopulationResponce, resPopError := linker.SendPostRequest(postURL, reserveRequestBody, *client)
		if resPopError != nil {
			linker.SendErrorAsJson("Error posting population response: ", resPopError, w)
			return
		} else {
			populationResponce = ReservePopulationResponce
		}
		defer ReservePopulationResponce.Body.Close()
	}

	defer populationResponce.Body.Close()

	//converting responce to json code
	convertedPopulation, convertingError := buildResponce(populationResponce, startY, endY)
	if convertingError != nil {
		linker.SendErrorAsJson("Error converting population response:", convertingError, w)
		return
	} else {
		json.NewEncoder(w).Encode(convertedPopulation)
	}

}

// building responce. Putting all response we need in structures and getting information we need
func buildResponce(response http.Response, startY int, endY int) (ResponsePopulation, error) {
	var populatinBuild CountryPopulationStructure
	decoder := json.NewDecoder(response.Body)
	if errDecoder1 := decoder.Decode(&populatinBuild); errDecoder1 != nil {
		log.Println("Error in decoding response from country: ", errDecoder1.Error())
	}

	var populationResponce ResponsePopulation
	for _, element := range populatinBuild.Data.Population {
		if startY == 0 && endY == 0 {
			populationResponce.Values = append(populationResponce.Values, element)
		} else if startY > 0 && endY == -1 {
			if startY == element.Year {
				populationResponce.Values = append(populationResponce.Values, element)
			}
		} else if startY > 0 && endY == 0 {
			if element.Year >= startY {
				populationResponce.Values = append(populationResponce.Values, element)
			}
		} else if endY > 0 && startY == 0 {
			if element.Year <= endY {
				populationResponce.Values = append(populationResponce.Values, element)
			}
		} else if endY >= startY {
			if element.Year >= startY && element.Year <= endY {
				populationResponce.Values = append(populationResponce.Values, element)
			}
		} else if endY < startY {
			fmt.Println(startY, endY)
			return populationResponce, errors.New("Start year cannot be less than end year")
		}
	}
	if len(populationResponce.Values) == 0 {
		return populationResponce, errors.New("No populations found for given years")
	}
	populationResponce.Mean = GetAvaragePopulation(populationResponce)
	return populationResponce, nil
}

func GetAvaragePopulation(population ResponsePopulation) int {
	if len(population.Values) > 0 {
		avarage := 0
		for _, element := range population.Values {
			avarage += element.Value
		}
		return avarage / len(population.Values)
	}
	return 0
}
