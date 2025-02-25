package utility

import (
	"fmt"
	"net/http"
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
