package maths

import (
	"time"
	"math/rand"
)

func GetRand(num int) int{
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(num)
	return rnd
}
func BetweenRand(numMin int , numMax int) int{
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(numMax - numMin) + numMin
	return rnd
}