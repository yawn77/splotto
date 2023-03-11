package generators

import (
	"testing"

	"github.com/yawn77/splotto/pkg/lotto/player"
)

func TestRandom(t *testing.T) {
	// arrange
	g := NewRandomGenerator()

	// act
	draw := g.GenerateNumbers(player.LottoHistory{})

	// assert
	if len(draw) != 6 {
		t.Fatal("expected 6 numbers to be drawn")
	}
	numberSet := make(map[int]bool)
	for i := 0; i < len(draw); i++ {
		n := draw[i]
		if n < 1 || n > 22 {
			t.Fatalf("%d is not between 0 and 22", i)
		}
		numberSet[n] = true
	}
	if len(numberSet) != 6 {
		t.Fatal("expected 6 different numbers to be drawn: %v", numberSet)
	}
}
