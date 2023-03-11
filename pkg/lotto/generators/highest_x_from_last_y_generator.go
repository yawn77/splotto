package generators

import (
	"math/rand"
	"sort"

	"github.com/gookit/slog"
	"github.com/yawn77/splotto/pkg/lotto"
)

type Histogram map[int]int

type HighestXfromLastYGenerator struct {
	numbersToKeep   int
	drawsToConsider int
}

func NewHighestXfromLastYGenerator(numbersToKeep int, drawsToConsider int) HighestXfromLastYGenerator {
	return HighestXfromLastYGenerator{numbersToKeep, drawsToConsider}
}

func min(numbers ...int) int {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}
	return min
}

func generateHistogram(history lotto.LottoHistory, drawsToConsider int) (histogram Histogram) {
	histogram = make(Histogram)
	for i := 0; i < min(len(history), drawsToConsider); i++ {
		numbers, err := history[i].Numbers()
		if err != nil {
			slog.Warn("will not consider broken draw for HighestXfromLastYGenerator, will try to use the next one instead")
			i += 1
			continue
		}
		for _, v := range numbers {
			_, ok := histogram[v]
			if ok {
				histogram[v] += 1
			} else {
				histogram[v] = 1
			}
		}
	}
	return histogram
}

func histogramToSortedList(histogram Histogram) (sortedNumbers []int) {
	sortedNumbers = make([]int, 0, len(histogram))
	for number := range histogram {
		sortedNumbers = append(sortedNumbers, number)
	}
	sort.SliceStable(sortedNumbers, func(i, j int) bool {
		return histogram[sortedNumbers[i]] > histogram[sortedNumbers[j]]
	})
	return sortedNumbers
}

func (g HighestXfromLastYGenerator) GenerateNumbers(history lotto.LottoHistory) (draw lotto.Draw, randomNumbers int) {
	draw = lotto.NewDraw()
	sortedNumbers := histogramToSortedList(generateHistogram(history, g.drawsToConsider))
	for i := 0; i < min(len(sortedNumbers), g.numbersToKeep, 6); i++ {
		draw.AddNumber(sortedNumbers[i])
	}
	randomNumbers = 6 - draw.Size()
	for !draw.IsValid() {
		n := rand.Intn(22) + 1
		ok := draw.AddNumber(n)
		if !ok {
			slog.Warn("tried to add invalid number (%d) to draw %v", n, draw)
		}
	}
	return draw, randomNumbers
}
