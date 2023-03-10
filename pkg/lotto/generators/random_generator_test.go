package generators

import (
	"testing"

	"github.com/yawn77/splotto/pkg/lotto/player"
)

func TestRandom(t *testing.T) {
	// arrange
	g := NewRandomGenerator()

	// act
	numbers := g.GenerateNumbers(player.LottoHistory{})

	// assert
	if len(numbers) != 6 {
		t.Fatal("expected 6 numbers to be generated")
	}
	set := make(map[int]bool)
	for i := 0; i < len(numbers); i++ {
		n := numbers[i]
		if n < 1 || n > 22 {
			t.Fatalf("%d is not between 0 and 22", i)
		}
		set[n] = true
	}
	if len(set) != 6 {
		t.Fatal("expected 6 different numbers to be generated")
	}
}
