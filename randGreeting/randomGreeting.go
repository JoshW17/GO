package main

import (
	"fmt"
	"math/rand"
)

func main() {
	greetings := [...]string{"Hello there!", "Howdy partner!", "Hola como estas", "Salutations good to be back on the air!"}
	numR := rand.Intn(3) + 1
	fmt.Println(greetings[numR])
}
