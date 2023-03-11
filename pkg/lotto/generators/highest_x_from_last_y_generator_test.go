package generators_test

import (
	"testing"

	"github.com/yawn77/splotto/pkg/lotto"
	"github.com/yawn77/splotto/pkg/lotto/generators"
)

func contains(s []int, n int) bool {
	for _, v := range s {
		if v == n {
			return true
		}
	}

	return false
}

func drawFromArray(numbers [6]int) lotto.Draw {
	draw := lotto.NewDraw()
	for _, n := range numbers {
		draw.AddNumber(n)
	}
	return draw
}

func TestGenerateNumbers(t *testing.T) {
	tests := []struct {
		testcase        string
		numbersToKeep   int
		drawsToConsider int
		history         lotto.LottoHistory
		expected        []int
	}{
		{
			"happy test case",
			6,
			1,
			lotto.LottoHistory{drawFromArray([6]int{1, 2, 3, 4, 5, 6})},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"more draws in history than considered",
			6,
			1,
			lotto.LottoHistory{drawFromArray([6]int{1, 2, 3, 4, 5, 6}), drawFromArray([6]int{1, 2, 3, 4, 5, 7}), drawFromArray([6]int{1, 2, 3, 4, 5, 7})},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"less draws in history than considered",
			6,
			2,
			lotto.LottoHistory{drawFromArray([6]int{1, 2, 3, 4, 5, 6})},
			[]int{1, 2, 3, 4, 5, 6},
		},
		{
			"empty history",
			6,
			1,
			lotto.LottoHistory{},
			[]int{},
		},
		{
			"keep more than 6 numbers",
			7,
			1,
			lotto.LottoHistory{drawFromArray([6]int{1, 2, 3, 4, 5, 7}), drawFromArray([6]int{1, 2, 3, 4, 8, 6}), drawFromArray([6]int{1, 2, 3, 9, 5, 6})},
			[]int{1, 2, 3},
		},
		{
			"keep less than 6 numbers",
			3,
			3,
			lotto.LottoHistory{drawFromArray([6]int{1, 2, 3, 4, 5, 6}), drawFromArray([6]int{1, 2, 7, 4, 8, 6}), drawFromArray([6]int{1, 2, 9, 4, 5, 10})},
			[]int{1, 2, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testcase, func(t *testing.T) {
			// arrange
			g := generators.NewHighestXfromLastYGenerator(tt.numbersToKeep, tt.drawsToConsider)

			// act
			draw, randomNumbers := g.GenerateNumbers(tt.history)

			// assert
			ok := draw.IsValid()
			if !ok {
				t.Fatal("draw is invalid", draw)
			}
			numbers, _ := draw.Numbers()
			for _, n := range tt.expected {
				if !contains(numbers[:], n) {
					t.Fatalf("expected %d to be in draw %v", n, numbers)
				}
			}
			random := 6 - tt.numbersToKeep
			if random < 0 {
				random = 0
			} else if len(tt.history) == 0 {
				random = 6
			}
			if randomNumbers != random {
				t.Fatalf("expected %d numbers to be random but %d numbers are random", random, randomNumbers)
			}
		})
	}
}
