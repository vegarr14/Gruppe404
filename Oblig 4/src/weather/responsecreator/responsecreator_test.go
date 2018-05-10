package responsecreator

import (
  "testing"
  "strings"
  "time"
  "weather/responsecreator"
)

func TestTimePositive(t *testing.T) {
  test := responsecreator.Time(0, 10000, 500)
  svar := "sunset in 2h38m20s"
  if strings.Compare(test, svar) != 0 {
    t.Errorf("test failed. got %d, want %d.", test, svar)
  }
}

func TestTimeNegative(t *testing.T) {
  test := responsecreator.Time(1525940000, 1525970000, time.Now().Unix())
  svar := "sunset in 4h4m48s"
  if strings.Compare(test, svar) != 0 {
    t.Errorf("test failed. got %d, want %d.", test, svar)
  }
}
