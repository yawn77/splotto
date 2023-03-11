package generators

import (
	"flag"
	"testing"

	"github.com/yawn77/splotto/pkg/lotto"
)

// add for successful tests
var Username = flag.String("username", "", "Username")
var Password = flag.String("password", "", "User password")

func TestRandom(t *testing.T) {
	// arrange
	g := NewRandomGenerator()

	// act
	draw := g.GenerateNumbers(lotto.LottoHistory{})

	// assert
	ok := draw.IsValid()
	if !ok {
		t.Fatal("draw is invalid", draw)
	}
}
