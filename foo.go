package foo

import "math/rand"

//SomeFunc does returns something - yay!
func SomeFunc(x int) int {
	min := 10
	max := 30
	return rand.Intn(max-min) + min
}
