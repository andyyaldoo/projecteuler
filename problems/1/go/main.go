package main

import (
	"log"
	"time"
)

func main() {
	defer timeTrack(time.Now(), "the opreation")
	multiples := []int{3, 5}
	total := 0
	for c := 1; c < 1000; c++ {
		for i := len(multiples) - 1; i >= 0; i-- {
			// log.Printf("try to divide %d by %d", c, multiples[i])
			if isMultipleOf(c, multiples[i]) {
				// log.Printf("it is multiple of 3 and 5")
				total += c
				// log.Printf("total is now %d", total)
				break
			}
		}
	}

	log.Printf("The answer is %d", total)
}

func isMultipleOf(n, d int) bool {
	return n%d == 0
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
