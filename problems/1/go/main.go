package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("The answer is %d", sumOfMultiplesOf3And5(1000))
}

func sumOfMultiplesOf3And5(max int) int {
	defer timeTrack(time.Now(), "the operation")
	total := 0
	for c := 1; c < max; c++ {
		if isMultipleOf3Or5(c) {
			// log.Printf("it is multiple of 3 and 5")
			total += c
			// log.Printf("total is now %d", total)
		}
	}
	return total
}

func isMultipleOf3Or5(n int) bool {
	return n%3 == 0 || n%5 == 0
}

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
