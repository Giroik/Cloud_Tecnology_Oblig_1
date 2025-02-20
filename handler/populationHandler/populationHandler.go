package populationHandler

import (
	"OBLIG_1/utility"
	"fmt"
	"net/http"
)

func PopulationHandler(w http.ResponseWriter, r *http.Request) {
	isoCode, startY, endY := utility.FormatISOandPopulationYears(r.URL.String())
	fmt.Println("isoCode:", isoCode)
	fmt.Println("startY:", startY)
	fmt.Println("endY:", endY)
}
