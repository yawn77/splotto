package lotto_test

import (
	"sort"
	"testing"

	"github.com/yawn77/splotto/pkg/lotto"
)

func TestNewDraw(t *testing.T) {
	// arrange

	// act
	d := lotto.NewDraw()
	for i := 1; i <= 6; i++ {
		ok := d.AddNumber(i)
		if !ok {
			t.Fatal("adding number to draw would make the draw invalid")
		}
	}

	// assert
	ok := d.IsValid()
	if !ok {
		t.Fatal("generated draw is invalid")
	}
}

func TestNumbers(t *testing.T) {
	// arrange
	d := lotto.NewDraw()
	for i := 1; i <= 6; i++ {
		ok := d.AddNumber(i)
		if !ok {
			t.Fatal("adding number to draw would make the draw invalid")
		}
	}

	// act
	numbers, err := d.Numbers()

	// assert
	if err != nil {
		t.Fatal(err)
	}
	expected := [6]int{1, 2, 3, 4, 5, 6}
	sort.Ints(numbers[:])
	if numbers != expected {

		t.Fatalf("expected %v == %v", numbers, expected)
	}
}

func TestInvalidNumbers(t *testing.T) {
	// arrange
	d := lotto.NewDraw()
	for i := 1; i <= 2; i++ {
		ok := d.AddNumber(i)
		if !ok {
			t.Fatal("adding number to draw would make the draw invalid")
		}
	}

	// act
	_, err := d.Numbers()

	// assert
	if err == nil {
		t.Fatal("expected error as len(draw) != 6")
	}
}

func TestInvalidAdd(t *testing.T) {
	// arrange
	d := lotto.NewDraw()

	// act & assert
	ok := d.AddNumber(0)
	if ok {
		t.Fatal("expected error as n < 1")
	}
	ok = d.AddNumber(23)
	if ok {
		t.Fatal("expected error as number > 22")
	}
	for i := 1; i <= 6; i++ {
		ok := d.AddNumber(i)
		if !ok {
			t.Fatal("adding number to draw would make the draw invalid")
		}
	}
	ok = d.AddNumber(22)
	if ok {
		t.Fatal("expected error as draw is already valid")
	}
}
