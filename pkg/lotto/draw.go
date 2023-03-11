package lotto

// errors
const (
	invalidDrawError = LottoError("draw is invalid")
)

func NewDraw() Draw {
	return Draw{make(map[int]interface{}, 6)}
}

type Draw struct {
	numbers map[int]interface{}
}

func (d Draw) Numbers() (numbers [6]int, err error) {
	ok := d.IsValid()
	if !ok {
		return numbers, invalidDrawError
	}
	keys := make([]int, 0, len(d.numbers))
	for k := range d.numbers {
		keys = append(keys, k)
	}
	return [6]int(keys), nil
}

func (d Draw) AddNumber(n int) bool {
	if d.IsValid() || n < 1 || n > 22 {
		return false
	}

	d.numbers[n] = nil
	return true
}

func (d Draw) IsValid() bool {
	l := len(d.numbers)
	if l != 6 {
		return false
	}
	for n := range d.numbers {
		if n < 1 || n > 22 {
			return false
		}
	}
	return true
}
