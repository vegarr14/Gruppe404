package responsecreator

import (
  "fmt"
  "time"
  "strconv"
)

func Getresponse(input int) (string, string) {
  var s string
  var imgPath string
  if 200<= input && input<=232  {
    s ="thunder"
    imgPath = "weather/icons/torden.png"
  }else if 300<= input && input<= 321 {
    s ="fo' drizzle muh nizzle"
    imgPath = "weather/icons/solmedlittskyer.png"
  }else if 500<= input && input <= 531 {
    s= "rain"
    imgPath = "weather/icons/regn.png"
  }else if 600<= input && input <= 622 {
    s= "you know nothing (jon)snow"
    imgPath = "weather/icons/sno.png"
  }else if 701<= input && input <= 781 {
    s= "there is something bad happening"
    imgPath = "weather/icons/littoveskyet.png"
  }else if 800 == input {
    s= "all clear"
    imgPath = "weather/icons/sol.png"
  }else if 801<= input && input <= 804 {
    s= "clouds"
    imgPath = "weather/icons/overskyet.png"
  }
  return s, imgPath
}

func Time(sunrise int64, sunset int64, unixnow int64) (string) {
  var t string
  var d time.Duration
  var err error

  if unixnow < sunrise {
    d,err = time.ParseDuration(strconv.FormatInt(sunrise - unixnow, 10) + "s")
    CheckError(err)
    t = "sunrise in " + d.String()
  } else if unixnow < sunset {
    d,err = time.ParseDuration(strconv.FormatInt(sunset - unixnow, 10) + "s")
    CheckError(err)
    t = "sunset in " + d.String()
  } else if unixnow > sunset {
    d,err = time.ParseDuration(strconv.FormatInt(unixnow - sunset, 10) + "s")
    CheckError(err)
    t = "sunset was " + d.String() + " ago"
  }
  return t
}

func CheckError(err error) {
  if err != nil {
    fmt.Println("error:", err)
  }
}
