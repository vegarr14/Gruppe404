package weather

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
)
type Feilmelding struct {
  Cod string
  Message string
}

type CWStruct struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base string `json:"base"`
	Main struct {
		Temp     float64 `json:"temp"`
		Pressure int     `json:"pressure"`
		Humidity int     `json:"humidity"`
		TempMin  float64 `json:"temp_min"`
		TempMax  float64 `json:"temp_max"`
	} `json:"main"`
	Visibility int `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Dt  int `json:"dt"`
	Sys struct {
		Type    int     `json:"type"`
		ID      int     `json:"id"`
		Message float64 `json:"message"`
		Country string  `json:"country"`
		Sunrise int     `json:"sunrise"`
		Sunset  int     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}

func CurrentWeather(location string) (Feilmelding, CWStruct) {
  var m CWStruct
  var feil Feilmelding
  link := "https://api.openweathermap.org/data/2.5/weather?q=" + location + "&appid=b1bf40e9707aee87cf7f39cd96df39b1"
  resp, err := http.Get(link)
  CheckError(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  _ = json.Unmarshal(body, &feil)
  _ = json.Unmarshal(body, &m)
  return feil,m
}

func CheckError(err error) {
  if err != nil {
    fmt.Println("error:", err)
  }
}
