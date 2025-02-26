package utility

/*
Creating structure for status information about API's
*/

type PopulationStructur struct {
	Mean   int `json:"mean"`
	Values []struct {
		Year  int `json:"year"`
		Value int `json:"value"`
	} `json:"values"`
}
type CountriesNowISOCountries struct {
	Error bool          `json:"error"`
	Msg   string        `json:"msg"`
	Data  []CountryData `json:"data"`
}

type CountryData struct {
	Name        string `json:"name"`
	OfisialName string `json:"ofisialName"`
	Iso2        string `json:"Iso2"`
	Iso3        string `json:"Iso3"`
}
type CountryIsoCheck struct {
	Name NameInfo `json:"name"`
	Cca2 string   `json:"cca2"`
	Cca3 string   `json:"cca3"`
}
type NameInfo struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type ErrorMassage struct {
	Error    bool   `json:"error"`
	ErrorMsg string `json:"msg"`
}
