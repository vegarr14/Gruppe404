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
    s= "You might want to take an umbrella with you because\nI'll be crying in the rain"
    imgPath = "weather/icons/regn.png"
  }else if 600<= input && input <= 622 {
    s= "You might need some gloves because the clouds decided to\nLet it snow, let it snow, let it snow"
    imgPath = "weather/icons/sno.png"
  }else if 701<= input && input <= 781 {
    s= "There is something bad happening here!\n What it is ain't exactly clear!"
    imgPath = "weather/icons/Tornado.png"
  }else if 800 == input {
    s= "Bring your sunglasses with you because\nYou are my sunshine, my only sunshine!"
    imgPath = "weather/icons/sol.png"
  }else if 801<= input && input <= 804 {
    s= "Take me to the clouds above"
    imgPath = "weather/icons/overskyet.png"
  }
  return s, imgPath
}

/*Funksjon som tar tiden for soloppgang, solnedgang og tiden i nå som input i unix time.
* Sammenlinger input verdiene og returnerer en string som skal vises på siden.
* Stringen viser hvor lenge det er til soloppgang eller solnedgang,
* eller hvor lenge det er siden solnedgang.
*/
func Time(sunrise int64, sunset int64, unixnow int64) (string) {
  var t string
  var d time.Duration
  var err error

  if unixnow < sunrise {
    d,err = time.ParseDuration(strconv.FormatInt(sunrise - unixnow, 10) + "s")
    CheckError(err)
    t = "Sunrise is in " + d.String()
  } else if unixnow < sunset {
    d,err = time.ParseDuration(strconv.FormatInt(sunset - unixnow, 10) + "s")
    CheckError(err)
    t = "Sunset is in " + d.String()
  } else if unixnow > sunset {
    d,err = time.ParseDuration(strconv.FormatInt(unixnow - sunset, 10) + "s")
    CheckError(err)
    t = "Sunset was " + d.String() + " ago"
  }
  return t
}

func CheckError(err error) {
  if err != nil {
    fmt.Println("error:", err)
  }
}
