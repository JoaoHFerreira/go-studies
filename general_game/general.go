package main

import (
	"fmt"
	"math/rand"
)



type hand struct {
	playTimes		int
    availablePlays  map[string]int
}

func newHand() hand {
	 availablePlays := map[string]int{
		"FullHouse":  0,
		"fiveOfAKind": 0,
		"fourOfAKind": 0,
		"follow": 0,
		"one": 0,
		"two": 0,
		"three": 0,
		"four": 0,
		"five": 0,
		"six": 0,
	   }
	   
	return hand{3, availablePlays}
}

func (h hand) GetUnitDiece() int{
	return rand.Intn(6) + 1
}

func (h hand) GetGeneralHand() []int{
	const plays = 3
	arrayPlays := make([]int, plays)

	for i := 0; i < plays; i++ {
		arrayPlays[i] = h.GetUnitDiece()
	}

	return arrayPlays
}

func main(){
	fmt.Println("Hello World")
}