package generators

import (
	"math/rand"

	"github.com/yawn77/splotto/pkg/lotto/player"
)

type RandomGenerator struct{}

func NewRandomGenerator() RandomGenerator {
	return RandomGenerator{}
}

func (RandomGenerator) GenerateNumbers(history player.LottoHistory) (draw player.Draw) {
	numberSet := make(map[int]bool)
	for len(numberSet) < 6 {
		numberSet[rand.Intn(22)+1] = true
	}
	i := 0
	for number := range numberSet {
		draw[i] = number
		i++
	}
	return draw
}
