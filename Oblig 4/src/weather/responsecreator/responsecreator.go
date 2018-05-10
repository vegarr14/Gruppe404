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
    s ="Thunder! Nananananananan Thunder!"
    imgPath = "weather/icons/torden.png"
  }else if 300<= input && input<= 321 {
    s ="Fo' drizzle muh nizzle"
    imgPath = "weather/icons/solmedlittskyer.png"
  }else if 500<= input && input <= 531 {
    s= "I'll be crying in the rain"
    imgPath = "weather/icons/regn.png"
  }else if 600<= input && input <= 622 {
    s= "Let it snow, let it snow, let it snow"
    imgPath = "weather/icons/sno.png"
  }else if 701<= input && input <= 781 {
    s= "There is something bad happening here!\n What it is ain't exactly clear!"
    imgPath = "weather/icons/Tornado.png"
  }else if 800 == input {
    s= "You are my sunshine, my only sunshine!"
    imgPath = "weather/icons/sol.png"
  }else if 801<= input && input <= 804 {
    s= "Take me to the clouds above"
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
