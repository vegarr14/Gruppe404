package weather

import (
  "weather"
  "fmt"
  "testing"
)
//Sjekker om to verdi i structen er tomme.
func TestWeatherPositive(t *testing.T) {
  data := weather.Weather("oslo")
  if data.Current.Base == "" {
    t.Errorf("Test failed. Struct is empty?")
  } else if data.Forecast.List[1].Time == "" {
    t.Errorf("Test failed. unix time not converted to dates")
  }
  if t.Failed() == false {
    fmt.Println(t.Name() + " passed")
  }
}

//Tester for en by som skal gi feilmelding
func TestWeatherNegative(t *testing.T) {
  data := weather.Weather("invalid city")
  fmt.Println(data)
  if data.Feilmelding.Message == "" {
    t.Errorf("Test failed. Invalid location was used. There should be an error message")
  }
  if t.Failed() == false {
    fmt.Println(t.Name() + " passed")
  }
}
