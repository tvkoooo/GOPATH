package maths

import (
	"math/rand"
	"time"
)

func GetRand(num int) int {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(num)
	return rnd
}
func BetweenRand(numMin int, numMax int) int {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(numMax-numMin) + numMin
	return rnd
}
