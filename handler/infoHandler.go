package handler

import "net/http"

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	output := "Request: " + LINEBREAK
	output += "URL Path: " + r.URL.Path + LINEBREAK
	output += "Path value: " + r.PathValue("Country_Code") + LINEBREAK
	output += "Path value: " + r.PathValue("City_limit") + LINEBREAK
	output += "Method: " + r.Method + LINEBREAK

}
