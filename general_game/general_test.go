package main

import (
	"testing"
	"fmt"
)

func TestGetUnitDiece(t *testing.T) {
	hand := newHand()
	result := hand.GetUnitDiece()
	expectedMaxLimit := 6
	expectedMinLimit := 1

	if result > expectedMaxLimit || result < expectedMinLimit {
		fmt.Println(result <= expectedMaxLimit)
		fmt.Println(result >= expectedMinLimit)
		t.Errorf("Value should be between %d and %d\nReceived %d", expectedMaxLimit, expectedMinLimit ,result)
	}
}

func TestGeneralHand(t *testing.T) {
	const plays = 3
	hand := newHand()
	result := hand.GetGeneralHand()
	if len(result) != plays {
		t.Errorf("Expected %d plays, got %d", plays, len(result))
	}

	for _, play := range result {
		if play > 6 || play < 1 {
			t.Errorf("Value should be between %d and %d\nReceived %d", 1, 6, play)
		}
	}
}

func simplePrint(t *testing.T){
	hand := newHand()
	if 1 < 3 {
		t.Errorf("%d", hand.GetGeneralHand())
	}
	
}