package weather

import (
  "fmt"
  "net/http"
  "io/ioutil"
  "encoding/json"
  "weather/responsecreator"
  "time"
  "strconv"
)
var ImgPath string

func CheckError(err error) {
  if err != nil {
    fmt.Println("error:", err)
  }
}

type WStruct struct {
Feilmelding struct{
  Cod string
  Message string
  }  
Current struct {
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
		Sunrise int64     `json:"sunrise"`
		Sunset  int64     `json:"sunset"`
	} `json:"sys"`
	ID   int    `json:"id"`
	Name string `json:"name"`
	Cod  int    `json:"cod"`
}
Forecast struct {
  Cod     string  `json:"cod"`
	Message float64 `json:"message"`
	Cnt     int     `json:"cnt"`
	List    []struct {
    Time string
		Dt   int64 `json:"dt"`
		Main struct {
			Temp      float64 `json:"temp"`
			TempMin   float64 `json:"temp_min"`
			TempMax   float64 `json:"temp_max"`
			Pressure  float64 `json:"pressure"`
			SeaLevel  float64 `json:"sea_level"`
			GrndLevel float64 `json:"grnd_level"`
			Humidity  int     `json:"humidity"`
			TempKf    float64 `json:"temp_kf"`
		} `json:"main"`
		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Wind struct {
			Speed float64 `json:"speed"`
			Deg   float64 `json:"deg"`
		} `json:"wind"`
		Rain struct {
			ThreeH float64 `json:"3h"`
		} `json:"rain"`
		Sys struct {
			Pod string `json:"pod"`
		} `json:"sys"`
		DtTxt string `json:"dt_txt"`
		Snow  struct {
			ThreeH float64 `json:"3h"`
		} `json:"snow,omitempty"`
	} `json:"list"`
	City struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Coord struct {
			Lat float64 `json:"lat"`
			Lon float64 `json:"lon"`
		} `json:"coord"`
		Country    string `json:"country"`
		Population int    `json:"population"`
	} `json:"city"`
}
  //Modified Data
  Response string
  Time string
  ImgPath string
}

func Weather(location string) (Feilmelding, WStruct) {
  var m WStruct
  link := "https://api.openweathermap.org/data/2.5/weather?q=" + location + "&units=metric&appid=b1bf40e9707aee87cf7f39cd96df39b1"
  body := getData(link)
  _ = json.Unmarshal(body, &m.Feilmelding)
  err2 := json.Unmarshal(body, &m.Current)
  if err2 != nil {
    return m
  }
  link2 := "https://api.openweathermap.org/data/2.5/forecast?q=" + location + "&units=metric&appid=b1bf40e9707aee87cf7f39cd96df39b1"
  body2 := getData(link2)
  _ = json.Unmarshal(body2, &m.Forecast)
  m.Response, ImgPath = responsecreator.Getresponse(m.Current.Weather[0].ID)
  m.Time = responsecreator.Time(m.Current.Sys.Sunrise, m.Current.Sys.Sunset, time.Now().Unix())

  for i := 0 ; i < len(m.Forecast.List) ; i++ {
    unix := m.Forecast.List[i].Dt
    time := time.Unix(unix, 0)
    dato := strconv.Itoa(time.Day()) + ". " + time.Month().String() + " " + strconv.Itoa(time.Hour()) + ":00"
    fmt.Println(unix, dato)
    m.Forecast.List[i].Time = dato
  }

  return m
}

func getData(link string) ([]byte) {
  resp, err := http.Get(link)
  CheckError(err)
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  return body
}
