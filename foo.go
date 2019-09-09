package foo

import (
  "math/rnd"
  )

//SomeFunction does returns something
func SomeFunc(x int) int {
  min := 10
  max := 30
  return rand.Intn(max - min) + min
}
