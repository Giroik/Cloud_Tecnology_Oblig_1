package constants

// urls
const COUNTRIES_NOW_API = "http://129.241.150.113:3500/api/v0.1"
const REST_COUNTRIES_API = "http://129.241.150.113:8080/v3.1"

// requests
const ENDPOINTALL = "/countries/iso/"
const ENDPOINTCOUNTRIES = "/all/"
const ENDPOINT_ALPHA = "/alpha/"
const ENDPOINT_CITIES = "/countries/cities"
const ENDPOINT_COUNT = "?limit=10"

const FRONT_PAGE = "/countryinfo/v1/"
const CONTRY_INFORMATION = "/countryinfo/v1/info/"
const CONTRY_POPULATION = "/countryinfo/v1/population/"
const CONTRY_STATUS = "/countryinfo/v1/status/"

const REQUEST_REST_COUNTRIES_API_ALPHA = REST_COUNTRIES_API + ENDPOINT_ALPHA
const COUNTRUES_NOW_ALL_CITIES = COUNTRIES_NOW_API + ENDPOINT_CITIES
const REQUEST_ISO_CODES = COUNTRIES_NOW_API + ENDPOINTALL
