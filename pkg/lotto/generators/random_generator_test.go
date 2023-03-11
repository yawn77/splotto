package generators_test

import (
	"flag"
	"testing"

	"github.com/yawn77/splotto/pkg/lotto"
	"github.com/yawn77/splotto/pkg/lotto/generators"
)

// add for successful tests
var Username = flag.String("username", "", "Username")
var Password = flag.String("password", "", "User password")

func TestRandom(t *testing.T) {
	// arrange
	g := generators.NewRandomGenerator()

	// act
	draw, _ := g.GenerateNumbers(lotto.LottoHistory{})

	// assert
	ok := draw.IsValid()
	if !ok {
		t.Fatal("draw is invalid", draw)
	}
}
