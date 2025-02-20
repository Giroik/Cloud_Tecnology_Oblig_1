package countryInfoHandler

// Structure for country
type CountryInfoStructure struct {
	Name       NameInfo          `json:"name"`
	Continents []string          `json:"continents"`
	Population int               `json:"population"`
	Languages  map[string]string `json:"languages"`
	Borders    []string          `json:"borders"`
	Flags      FlagsInfo         `json:"flags"`
	Capital    []string          `json:"capital"`
	Cities     []string          `json:"cities"`
}

type CitiesInfoStructure struct {
	Error  bool     `json:"error"`
	Msg    string   `json:"msg"`
	Cities []string `json:"data"`
}

type NameInfo struct {
	Common   string `json:"common"`
	Official string `json:"official"`
}

type FlagsInfo struct {
	Png string `json:"png"`
	Svg string `json:"svg"`
}

type CityData struct {
	City    string `json:"city"`
	Country string `json:"country"`
}
