package handler

import (
	"fmt"
	"log"
	"net/http"
)

const LINEBREAK = "\n"

func FrontPageHandler(w http.ResponseWriter, r *http.Request) {
	output := "Hello and welcome to api for country information! \n" +
		"\nInformation: " + "This is a front page of application." +
		"\nHere can you get information and population about many different countrys\n" +
		"\nYou can choose between 3 different services.\n" +
		"1) countryinfo/v1/status/ \n" +
		"2) countryinfo/v1/info \n" +
		"3) countryinfo/v1/population \n" +
		"Below you will find information about how to choose country you want get info about\n\n\n" +
		"STATUS: \nTo get status of all api's we use type: " + "../countryinfo/v1/status/ \n\n\n" +
		"INFO:\nTo get information about county you have to use countries 2/3 letter code. This codes you can find in link bellow:\n" +
		"https://en.wikipedia.org/wiki/List_of_ISO_3166_country_codes" + "\n\n" +
		"After you found 2/3 letter code you can write it in link: /countryinfo/v1/info/{code}\n" +
		"For example '/countryinfo/v1/info/no' will return information about Norway\n" +
		"If you add '/countryinfo/v1/info/no?limit=20'  it will return information about Norway with list of 20 cities (default value 10)\n\n\n" +
		"POPULATION:\nTo get population of any country you have to use same countries 2/3 letter code as in INFO.\n" +
		"To get get population you have to use link: 'countryinfo/v1/population/{code}' \n" +
		"It will list population of country of all time and average population for all years \n" +
		"\nTo limit years use '../population/{code}?limit={startYear}-{EndYear}'\n" +
		"1) You can use '?limit=2000' to list population in country in 2000 \n" +
		"2) You can use '?limit=2000-' to list population in county after 2000\n" +
		"3) You can use '?limit=-2000' to list population in county before 2000\n"

	_, err := fmt.Fprintf(w, "%v", output)
	if err != nil {
		log.Println("An error occurred: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
