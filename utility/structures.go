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

type PopulationStructur struct {
	Mean   int `json:"mean"`
	Values []struct {
		Year  int `json:"year"`
		Value int `json:"value"`
	} `json:"values"`
}
