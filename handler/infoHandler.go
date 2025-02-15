package handler

import (
	"OBLIG_1/constants"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func InfoHandler(w http.ResponseWriter, request *http.Request) {
	// Create new request with new url
	newUrl := constants.REST_COUNTRIES_API +
		constants.ENDPONT_ALPHA +
		strings.TrimPrefix(request.URL.Path, constants.CONTRY_INFORMATION) ///countryinfo/v1/info/ru/---/countryinfo/v1/info/

	r, err := http.NewRequest(http.MethodGet, newUrl, nil)
	if err != nil {
		fmt.Errorf("Error in creating request:", err.Error())
	}

	// Setting content type -> effect depends on the service provider
	r.Header.Add("content-type", "application/json")

	// Instantiate the client
	client := &http.Client{}
	defer client.CloseIdleConnections()
	// Issue request
	res, err := client.Do(r)
	if err != nil {
		fmt.Errorf("Error in response:", err.Error())
	} else if res.StatusCode != http.StatusOK {
		fmt.Errorf("Error in response:", res.Status)
	} else if res.Header.Get("content-type") != "application/json" {
		fmt.Errorf("Header structure is not application/json ", res.Status)
	}

	fmt.Println("Status:", res.Status)
	fmt.Println("Status code:", res.StatusCode)

	fmt.Println("Content type:", res.Header.Get("content-type"))
	fmt.Println("Protocol:", res.Proto)

	output, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Errorf("Error when reading response: ", err.Error())
	}

	fmt.Println(string(output))
	/*
		// Decoding JSON
		decoder := json.NewDecoder(res.Body)
		var mp Joke
		if err := decoder.Decode(&mp); err != nil {
			log.Fatal(err)
		}

		// Printing decoded output
		fmt.Println(mp)
	*/
}
