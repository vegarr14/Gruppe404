package responsecreator

import (
  "testing"
  "strings"
  "time"
  "weather/responsecreator"
)

//Test som bruker forhåndssatte verdier og sjekker om det stemmer med svar stringen. 
func TestTimePositive(t *testing.T) {
  test := responsecreator.Time(0, 10000, 500)
  svar := "sunset in 2h38m20s"
  if strings.Compare(test, svar) != 0 {
    t.Errorf("test failed. got %d, want %d.", test, svar)
  }
}

//Test som skal feile siden time.now.unix gir en ny verdi hvert sekund.
func TestTimeNegative(t *testing.T) {
  test := responsecreator.Time(1525940000, 1525970000, time.Now().Unix())
  svar := "sunset in 4h4m48s"
  if strings.Compare(test, svar) != 0 {
    t.Errorf("test failed. got %d, want %d.", test, svar)
  }
}
