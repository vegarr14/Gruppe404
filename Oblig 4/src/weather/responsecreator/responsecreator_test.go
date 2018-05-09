package responsecreator_test

import (
  "testing"
  "responsecreator"
  "fmt"
)

func TestSomething(t *testing.T) {
  meme := responsecreator.Time(0, 10000, 500)
  fmt.Println(meme)
}
