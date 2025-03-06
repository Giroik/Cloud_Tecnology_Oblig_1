# Country Information Service

## Overview
This project is a REST web service built in Go that provides country-related information, including general country details, population data, and used API's status. The service fetches data from two external APIs:
- **CountriesNow API**: Is used to provides city population data.
- **REST Countries API**: Is used to provides general country information.

Service provides 3 main functions:
1. **Country Info**: Returns general country information.
2. **Country Population**: Returns population levels for given time frames.
3. **Servers status**: Provides a status overview of the services.

---

## API Endpoints

### 1. Country Population Endpoint
**Path:** `/countryinfo/v1/population/{:two_letter_country_code}{?limit={:startYear-endYear}}`

#### Request
- **Method:** `GET`
- **Path Parameters:**
  - `two_letter_country_code` (required): The ISO 3166-2 code for the country.
  - `limit` (optional): Filter results between `startYear` and `endYear`.

#### Example Request
```
population between 2005 and 2010
http://localhost:8080/countryinfo/v1/population/lv?limit=2005-2010

population before 2010
http://localhost:8080/countryinfo/v1/population/lv?limit=-2010

population after 2010
http://localhost:8080/countryinfo/v1/population/lv?limit=2010-


```

#### Response (JSON Example)
```json
{
    "country":"Latvia",
    "mean":2179004,
    "values":[
        {"year":2005,"value":2238799},
        {"year":2006,"value":2218357},
        {"year":2007,"value":2200325},
        {"year":2008,"value":2177322},
        {"year":2009,"value":2141669},
        {"year":2010,"value":2097555}]
}
```

---


### 2. Country Info Endpoint

**Path:** `/countryinfo/v1/info/{:two_letter_country_code}{?limit=10}`

#### Request
- **Method:** `GET`
- **Path Parameters:**
  - `two_letter_country_code` (required): The ISO 3166-2 code for the country.
  - `limit` (optional): Number of cities to return (sorted alphabetically).

#### Example Request
```
http://localhost:8080/countryinfo/v1/info/lv
```

#### Response (JSON Example)
```json
{
    "name":{"common":"Latvia","official":"Republic of Latvia"},
    "continents":["Europe"],
    "population":1901548,
    "languages":{"lav":"Latvian"},
    "borders":["BLR","EST","LTU","RUS"],
    "flags":{"png":"https://flagcdn.com/w320/lv.png",
    "svg":"https://flagcdn.com/lv.svg"},
    "capital":["Riga"],
    "cities":["Adazi","Agenskalns","Aizkraukle","Aizpute","Baldone","Balvi","Bauska","Brankas","Carnikava","Centrs"]
}
```

---


### 3. Diagnostics Endpoint
**Path:** `/countryinfo/v1/status/`

#### Request
- **Method:** `GET`

#### Example Request
```
http://localhost:8080/countryinfo/v1/status/
```

#### Response (JSON Example)
```json
{
    "countriesnowapi": 200,
    "restcountriesapi": 200,
    "version": "v1",
    "uptime": 3600
}
```

---


### Installation
1. Clone the repository:
   ```
   git clone https://github.com/Giroik/Cloud_Tecnology_Oblig_1.git
   cd Cloud_Tecnology_Oblig_1
   ```
2. Run the server:
   ```
   go run cmd/main.go
   ```
3. Open link in browser
    ```
    http://localhost:8080/countryinfo/v1/
    ```
4. Welcome to API, all information about links you will find on this address

---

## Code Structure
```
OBLIG 1
│── constants
│   └── constats.go
│── handlers
│   ├── countryInfoHandler
│   │   ├── infoHandler.go
│   │   └── models.go
│   ├── linker
│   │   └── linker.go
│   ├── populationHandler
│   │   ├── models.go
│   │   └── populationHandler.go
│   ├── statusHandler
│   │   ├── models.go
│   │   └── statusHandler.go
│   └── frontPageHandler.go
│── router
│   └── router.go
│── utility
│   ├── isoManipulations.go
│   └── structures.go
│── go.mod
|── main.go
│── README.md                 # Documentation
```

---
## Dev Features
1) In isoManipulations file i implemented 2 methods GetCountryNameByISO and GetReserveCountryNameByISO. First function is 
checking if this iso code exist in CountriesNow API if not it check also iso code in REST Countries API.
This methods also returns common name and official name what give more chances to find population and info by using both.
