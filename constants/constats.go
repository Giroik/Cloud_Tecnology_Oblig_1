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
const ENDPOINT_ALPHA_WITH_LIMIT = "/alpha?codes=1,"

const FRONT_PAGE = "/countryinfo/v1/"
const CONTRY_INFORMATION = "/countryinfo/v1/info/"
const CONTRY_POPULATION = "/countryinfo/v1/population/"
const CONTRY_STATUS = "/countryinfo/v1/status/"
const COUNTY_POPULATION_FILTER = "/countries/population"

const REQUEST_REST_COUNTRIES_API_ALPHA = REST_COUNTRIES_API + ENDPOINT_ALPHA
const COUNTRUES_NOW_ALL_CITIES = COUNTRIES_NOW_API + ENDPOINT_CITIES
const RESERVE_REQUEST_ISO_CODES = COUNTRIES_NOW_API + ENDPOINTALL
const REQUEST_ISO_CODE = REST_COUNTRIES_API + ENDPOINT_ALPHA
const REQUEST_FILTERD_POPULATION = COUNTRIES_NOW_API + COUNTY_POPULATION_FILTER

const NOT_FOUND = "notFound"

const RESET_LIMIT = 10

// INFORMATION PAGE
const INFO_CONST = "../info/{two_letter_country_code}{Limit_of_cities}/"
const POPULATION_CONST = "../population/{two_letter_country_code}"
const STATUS_CONST = "../status/"
