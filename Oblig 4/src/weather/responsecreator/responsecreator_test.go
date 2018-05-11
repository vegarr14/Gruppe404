package responsecreator

import (
  "testing"
  "strings"
  "time"
  "fmt"
  "weather/responsecreator"
)

//Test som bruker forhåndssatte verdier og sjekker om det stemmer med svar stringen.
func TestTimePositive(t *testing.T) {
  test := responsecreator.Time(0, 10000, 500)
  svar := "Sunset is in 2h38m20s"
  if strings.Compare(test, svar) != 0 {
    t.Errorf("test failed. got %d, want %d.", test, svar)
  }
  if t.Failed() == false {
    fmt.Println(t.Name() + " passed")
  }
}

//Test som skal feile siden time.now.unix gir en ny verdi hvert sekund.
func TestTimeNegative(t *testing.T) {
  test := responsecreator.Time(1525940000, 1525970000, time.Now().Unix())
  svar := "Sunset is in 4h4m48s"
  if strings.Compare(test, svar) == 0 {
    t.Errorf("test failed. got %d, want %d.", test, svar)
  }
  if t.Failed() == false {
    fmt.Println(t.Name() + " passed")
  }
}

//Test som bruker en kjent verdi
func TestGetresponsePositive(t *testing.T) {
  a, b := responsecreator.Getresponse(200)
  if a == "" && b == "" {
    t.Errorf("Test failed. Strings should not be empty")
  }
  if t.Failed() == false {
    fmt.Println(t.Name() + " passed")
  }
}

//Test som bruker en invalid verdi
func TestGetresponseNegative(t *testing.T) {
  a, _ := responsecreator.Getresponse(24)
  if strings.Contains(a, "Invalid") == false {
    t.Errorf("Test failed. value should be invalid.")
  }
  if t.Failed() == false {
    fmt.Println(t.Name() + " passed")
  }
}
