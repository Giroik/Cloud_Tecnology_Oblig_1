package handler

import (
	"OBLIG_1/constants"
	"fmt"
	"log"
	"net/http"
)

const LINEBREAK = "\n"

func FrontPageHandler(w http.ResponseWriter, r *http.Request) {
	output := "Hello and welcome to api for country information!" + LINEBREAK
	output += LINEBREAK + "Information: " + "This is a front page of application. To get some values use this requests: " + LINEBREAK + LINEBREAK
	output += constants.INFO_CONST + LINEBREAK
	output += "'/countryinfo/v1/info/no?limit=10' this page il provide you with information about country (in Norway) and display 10 of it's cities" + LINEBREAK + LINEBREAK
	output += constants.POPULATION_CONST + LINEBREAK
	output += "'/countryinfo/v1/info/population/no?limit=2000-2005' this page will provide you with information about population in country (Norway) in period from 2000 to 2005" + LINEBREAK
	output += "You also can use ?limit to show population at 1 year '?limit=2000' before given year '?limit=-2005' and after given year '?limit=2000-' " + LINEBREAK + LINEBREAK
	output += constants.STATUS_CONST + LINEBREAK
	output += "This page will provide you with information about apies status and program version" + LINEBREAK
	/*
		output += LINEBREAK + "Handlers: " + LINEBREAK
		for k, v := range r.Header {
			for _, vv := range v {
				output += k + ": " + vv + LINEBREAK
			}
		}*/

	_, err := fmt.Fprintf(w, "%v", output)
	if err != nil {
		log.Println("An error occurred: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
