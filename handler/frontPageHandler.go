package handler

import (
	"fmt"
	"log"
	"net/http"
)

const LINEBREAK = "\n"
const INFO_CONST = "../info/{two_letter_country_code}{Limit_of_cities}/"
const POPULATION_CONST = "../population/{two_letter_country_code}"
const STATUS_CONST = "../status/"

func FrontPageHandler(w http.ResponseWriter, r *http.Request) {
	output := "Request: " + LINEBREAK
	output += "URL Path: " + r.URL.Path + LINEBREAK
	output += "Method: " + r.Method + LINEBREAK
	output += LINEBREAK + "Information: " + "This is a front page of application. To get some values use this requests: " + LINEBREAK + LINEBREAK
	output += INFO_CONST + LINEBREAK
	output += POPULATION_CONST + LINEBREAK
	output += STATUS_CONST + LINEBREAK

	output += LINEBREAK + "Handlers: " + LINEBREAK
	for k, v := range r.Header {
		for _, vv := range v {
			output += k + ": " + vv + LINEBREAK
		}
	}

	_, err := fmt.Fprintf(w, "%v", output)
	if err != nil {
		log.Println("An error occurred: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
