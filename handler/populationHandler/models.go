package populationHandler

type CountryPopulationStructure struct {
	Error bool       `json:"error"`
	Msg   string     `json:"msg"`
	Data  DataStruct `json:"data"`
}

type DataStruct struct {
	Country    string      `json:"country"`
	Population []PopStruct `json:"populationCounts"`
}

type PopStruct struct {
	Year  int `json:"year"`
	Value int `json:"value"`
}

type ResponsePopulation struct {
	Country string      `json:"country"`
	Mean    int         `json:"mean"`
	Values  []PopStruct `json:"values"`
}
