package utility

/*
Creating structure for status information about API's
*/
type StatusStructur struct {
	CountriesNowAPI  int    `json:"countriesnowapi"`
	RestCountriesAPI int    `json:"restcountriesapi"`
	Version          string `json:"version"`
	Uptime           int64  `json:"uptime"`
}

type CountryInfoStructure struct {
	Name       string   `json:"name"`
	Continents []string `json:"continents"`
	Population int      `json:"population"`
	Languages  struct {
		Nno string `json:"nno"`
		Nob string `json:"nob"`
		Smi string `json:"smi"`
	} `json:"languages"`

	Borders []string `json:"borders"`
	Flag    string   `json:"flag"`
	Capital string   `json:"capital"`
	Cities  []string `json:"cities"`
}
