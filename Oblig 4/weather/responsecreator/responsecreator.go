package responsecreator

import (
  "time"
  "strconv"
)

func Getresponse(input int) (string) {
  var s string
  if 200<= input && input<=232  {
    s ="thunder"
  }else if 300<= input && input<= 321 {
    s ="fo' drizzle muh nizzle"
  }else if 500<= input && input <= 531 {
    s= "rain"
  }else if 600<= input && input <= 622 {
    s= "you know nothing (jon)snow"
  }else if 701<= input && input <= 781 {
    s= "there is something bad happening"
  }else if 800 == input {
    s= "all clear"
  }else if 801<= input && input <= 804 {
    s= "clouds"
  }
  return s
}

func Time(sunrise int64, sunset int64) (string) {
  var t string
  now := time.Now()
  Unixnow := now.Unix()
  if Unixnow < sunrise {
    t = "sunrise in " + strconv.FormatInt(sunrise - Unixnow, 10) + " seconds"
  } else if Unixnow < sunset {
    t = "sunset in " + strconv.FormatInt(sunset - Unixnow, 10) + " seconds"
  } else if Unixnow > sunset {
    t = "sunset was " + strconv.FormatInt(Unixnow - sunset, 10) + " seconds ago"
  }
  return t
}
