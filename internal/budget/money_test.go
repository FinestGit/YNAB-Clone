package budget

import (
	"errors"
	"math"
	"testing"
)

func TestNewMoney_roundTrip(t *testing.T) {
	money, err := NewMoney(123)
	if err != nil {
		t.Fatal(err)
	}
	if money.Cents() != 123 {
		t.Fatalf("got %d != 123", money.Cents())
	}
}

func TestAddMoney_PositivePositive(t *testing.T) {
	leftSide, err := NewMoney(100)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(200)
	if err != nil {
		t.Fatal(err)
	}
	sum, err := leftSide.Add(rightSide)
	if err != nil {
		t.Fatal(err)
	}
	if sum.Cents() != 300 {
		t.Fatalf("got %d != 300", sum.Cents())
	}
}

func TestAddMoney_NegativeNegative(t *testing.T) {
	leftSide, err := NewMoney(-100)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(-50)
	if err != nil {
		t.Fatal(err)
	}
	sum, err := leftSide.Add(rightSide)
	if err != nil {
		t.Fatal(err)
	}
	if sum.Cents() != -150 {
		t.Fatalf("got %d != -150", sum.Cents())
	}
}

func TestAddMoney_PositiveNegative(t *testing.T) {
	leftSide, err := NewMoney(100)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(-30)
	if err != nil {
		t.Fatal(err)
	}
	sum, err := leftSide.Add(rightSide)
	if err != nil {
		t.Fatal(err)
	}
	if sum.Cents() != 70 {
		t.Fatalf("got %d != 70", sum.Cents())
	}
}

func TestAddMoney_Zero(t *testing.T) {
	leftSide, err := NewMoney(0)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(100)
	if err != nil {
		t.Fatal(err)
	}
	sum, err := leftSide.Add(rightSide)
	if err != nil {
		t.Fatal(err)
	}
	if sum.Cents() != 100 {
		t.Fatalf("got %d != 100", sum.Cents())
	}
}

func TestAddMoney_Overflow(t *testing.T) {
	leftSide, err := NewMoney(math.MaxInt64)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(1)
	if err != nil {
		t.Fatal(err)
	}
	sum, err := leftSide.Add(rightSide)
	if sum.Cents() != 0 {
		t.Fatalf("got %d, instead of ErrMoneyOverflow", sum.Cents())
	}
	if !errors.Is(err, ErrMoneyOverflow) {
		t.Fatalf("did not get ErrMoneyOverflow")
	}
}

func TestAddMoney_Underflow(t *testing.T) {
	leftSide, err := NewMoney(math.MinInt64)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(-1)
	if err != nil {
		t.Fatal(err)
	}
	sum, err := leftSide.Add(rightSide)
	if sum.Cents() != 0 {
		t.Fatalf("got %d, instead of ErrMoneyUnderflow", sum.Cents())
	}
	if !errors.Is(err, ErrMoneyUnderflow) {
		t.Fatalf("did not get ErrMoneyUnderflow")
	}
}

func TestSubMoney_PositivePositive(t *testing.T) {
	leftSide, err := NewMoney(200)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(100)
	if err != nil {
		t.Fatal(err)
	}
	difference, err := leftSide.Sub(rightSide)
	if err != nil {
		t.Fatal(err)
	}
	if difference.Cents() != 100 {
		t.Fatalf("%d != 100", difference.Cents())
	}
}

func TestSubMoney_NegativeNegative(t *testing.T) {
	leftSide, err := NewMoney(-150)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(-50)
	if err != nil {
		t.Fatal(err)
	}
	difference, err := leftSide.Sub(rightSide)
	if err != nil {
		t.Fatal(err)
	}
	if difference.Cents() != -100 {
		t.Fatalf("%d != -100", difference.Cents())
	}
}

func TestSubMoney_PositiveNegative(t *testing.T) {
	leftSide, err := NewMoney(100)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(-50)
	if err != nil {
		t.Fatal(err)
	}
	difference, err := leftSide.Sub(rightSide)
	if err != nil {
		t.Fatal(err)
	}
	if difference.Cents() != 150 {
		t.Fatalf("got %d != 150", difference.Cents())
	}
}

func TestSubMoney_Zero(t *testing.T) {
	leftSide, err := NewMoney(100)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(0)
	if err != nil {
		t.Fatal(err)
	}
	difference, err := leftSide.Sub(rightSide)
	if err != nil {
		t.Fatal(err)
	}
	if difference.Cents() != 100 {
		t.Fatalf("got %d != 100", difference.Cents())
	}
}

func TestSubMoney_Overflow(t *testing.T) {
	leftSide, err := NewMoney(math.MaxInt64)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(-1)
	if err != nil {
		t.Fatal(err)
	}
	difference, err := leftSide.Sub(rightSide)
	if difference.Cents() != 0 {
		t.Fatalf("got %d, instead of ErrMoneyOverflow", difference.Cents())
	}
	if !errors.Is(err, ErrMoneyOverflow) {
		t.Fatalf("did not get ErrMoneyOverflow")
	}
}

func TestSubMoney_Underflow(t *testing.T) {
	leftSide, err := NewMoney(math.MinInt64)
	if err != nil {
		t.Fatal(err)
	}
	rightSide, err := NewMoney(1)
	if err != nil {
		t.Fatal(err)
	}
	difference, err := leftSide.Sub(rightSide)
	if difference.Cents() != 0 {
		t.Fatalf("got %d, instead of ErrMoneyUnderflow", difference.Cents())
	}
	if !errors.Is(err, ErrMoneyUnderflow) {
		t.Fatalf("did not get ErrMoneyUnderflow")
	}
}
