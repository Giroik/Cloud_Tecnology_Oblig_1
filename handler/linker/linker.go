package linker

import (
	"OBLIG_1/utility"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

func GetAPIStatus(url string, endPoint string) int {
	resp, err := http.Head(url + endPoint)
	defer resp.Body.Close()

	if err != nil {
		fmt.Println("Error making HEAD request:", err)
		return resp.StatusCode
	} else if resp.StatusCode != http.StatusOK {
		fmt.Println("Error making HEAD request, status code is:", resp.StatusCode)
		return resp.StatusCode
	}

	return resp.StatusCode
}

func SendGetRequest(url string, client http.Client) (http.Response, error) {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error in creating request: ", err.Error())
		return http.Response{}, errors.New("Error in creating request for url " + url)
	}
	request.Header.Add("Content-Type", "application/json")

	responce, err := client.Do(request)
	if err != nil {
		fmt.Println("Error in sending request:", err.Error())
		return http.Response{}, errors.New("Error in sending request: " + url + " " + err.Error())
	} else if responce.StatusCode != http.StatusOK {
		fmt.Println("Error in response status:", responce.Status)
		return http.Response{}, errors.New("Error in response status: " + responce.Status + " for url " + url)
	} else if responce.Header.Get("content-type") != "application/json" {
		fmt.Println("Header structure is not application/json ", responce.Status)
		return http.Response{}, errors.New("Header structure is not application/json")
	}
	return *responce, nil
}

func SendPostRequest(url string, req []byte, client http.Client) (http.Response, error) {

	request, err := http.NewRequest(http.MethodPost, url, strings.NewReader(string(req)))
	if err != nil {
		fmt.Println("Error in creating request for :", err.Error())
		return http.Response{}, errors.New("Error in creating request for url " + url)
	}
	request.Header.Add("Content-Type", "application/json")

	responce, err := client.Do(request)
	if err != nil {
		fmt.Println("Error in sending request:", err.Error())
		return http.Response{}, errors.New("Error in sending request: " + url + " " + err.Error())
	} else if responce.StatusCode != http.StatusOK {
		fmt.Println("Error in response status:", responce.Status)
		return http.Response{}, errors.New("Error in response status: " + responce.Status + "; user input not found in api " + url)
	}
	return *responce, nil
}

func SendErrorAsJson(errorText string, err error, w http.ResponseWriter) {
	var errorMassage utility.ErrorMassage
	errorMassage.ErrorMsg = errorText + " " + err.Error()
	errorMassage.Error = true
	json.NewEncoder(w).Encode(errorMassage)
}
