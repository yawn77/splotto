package generators

import (
	"math/rand"

	"github.com/gookit/slog"
	"github.com/yawn77/splotto/pkg/lotto"
)

type RandomGenerator struct{}

func NewRandomGenerator() RandomGenerator {
	return RandomGenerator{}
}

func (RandomGenerator) GenerateNumbers(history lotto.LottoHistory) (draw lotto.Draw) {
	draw = lotto.NewDraw()
	for !draw.IsValid() {
		n := rand.Intn(22) + 1
		ok := draw.AddNumber(n)
		if !ok {
			slog.Warn("tried to add invalid number (%d) to draw %v", n, draw)
		}
	}
	return draw
}
