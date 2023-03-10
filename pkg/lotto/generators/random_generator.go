package generators

import (
	"math/rand"

	"github.com/yawn77/splotto/pkg/lotto/player"
)

type RandomGenerator struct{}

func NewRandomGenerator() RandomGenerator {
	return RandomGenerator{}
}

func (RandomGenerator) GenerateNumbers(history player.LottoHistory) (numbers [6]int) {
	set := make(map[int]bool)
	for len(set) < 6 {
		set[rand.Intn(22)+1] = true
	}
	i := 0
	for k := range set {
		numbers[i] = k
		i++
	}
	return numbers
}
