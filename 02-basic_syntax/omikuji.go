package main

import (
	"math/rand"
	"time"
)

func main() {
	t := time.Now().UnixNano()
	rand.Seed(t)
	switch rand.Intn(6) + 1 {
	case 6:
		println("Congratulations! You have 大吉.")
	case 4, 5:
		println("That's good! You have 中吉.")
	case 3, 2:
		println("OK, You have 吉.")
	case 1:
		println("You have 凶. Don't lose heart.")
	}
}
