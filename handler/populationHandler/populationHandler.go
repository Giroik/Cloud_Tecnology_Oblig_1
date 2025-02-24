package populationHandler

import (
	"OBLIG_1/constants"
	"OBLIG_1/utility"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	//creating url for postRequest
	postURL := constants.REQUEST_FILTERD_POPULATION

	//Preparing client for tasks
	client := &http.Client{}
	defer client.CloseIdleConnections()

	//extracting information from users URL
	isoCode, startY, endY := utility.FormatISOandPopulationYears(r.URL.String())

	//extracting name to country throw ISO code
	countryName, countryError := utility.GetCountryNameByISO(isoCode)
	if countryError != nil {
		fmt.Fprintf(w, countryError.Error(), http.StatusInternalServerError)
	}
	fmt.Println("isoCode:", isoCode)
	fmt.Println("startY:", startY)
	fmt.Println("endY:", endY)
	fmt.Println("countryName:", countryName)

	requestBody, err := json.Marshal(map[string]string{"country": countryName})
	if err != nil {
		log.Println(err)
	}

	populationRequest, poperr := http.NewRequest("POST", postURL, strings.NewReader(string(requestBody)))
	if poperr != nil {
		fmt.Println("Error in creating request for Countries cities:", poperr.Error())
	}
	populationRequest.Header.Add("Content-Type", "application/json")

	populationResponce, err := client.Do(populationRequest)
	if err != nil {
		fmt.Println("Error in sending request to Countries cities:", err.Error())
	} else if populationResponce.StatusCode != http.StatusOK {
		fmt.Println("Error in response:", populationResponce.Status)
	} else if populationResponce.Header.Get("content-type") != "application/json" {
		fmt.Println("Header structure is not application/json ", populationResponce.Status)
	}

	defer populationResponce.Body.Close()
	w.Header().Set("Content-Type", "application/json")

	convertedPopulation, convertingError := buildResponce(*populationResponce, startY, endY)
	if convertingError != nil {
		fmt.Println("Error in converting population:", convertingError.Error())
		fmt.Fprintf(w, "Error: %s", convertingError.Error())

	} else {
		json.NewEncoder(w).Encode(convertedPopulation)
	}

}

func buildResponce(response http.Response, startY int, endY int) (ResponsePopulation, error) {
	var convertingError error = nil
	var populatinBuild CountryPopulationStructure
	decoder := json.NewDecoder(response.Body)
	if errDecoder1 := decoder.Decode(&populatinBuild); errDecoder1 != nil {
		log.Println("Error in decoding response from country: ", errDecoder1.Error())
	}

	var populationResponce ResponsePopulation
	for _, element := range populatinBuild.Data.Population {
		if startY == 0 && endY == 0 {
			populationResponce.Values = append(populationResponce.Values, element)
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
			return populationResponce, errors.New("Start year cannot be less than end year")
		}
	}

	populationResponce.Mean = GetAvaragePopulation(populationResponce)
	return populationResponce, convertingError
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
