package budget

import (
	"errors"
	"math"
)

type Money int64

var ErrMoneyOverflow = errors.New("money: integer overflow")
var ErrMoneyUnderflow = errors.New("money: integer underflow")

func NewMoney(cents int64) (Money, error) {
	return Money(cents), nil
}

func (money Money) Cents() int64 {
	return int64(money)
}

func (money Money) Add(rightSide Money) (Money, error) {
	if rightSide.Cents() > 0 && money.Cents() > math.MaxInt64-rightSide.Cents() {
		return 0, ErrMoneyOverflow
	}
	if rightSide.Cents() < 0 && money.Cents() < math.MinInt64-rightSide.Cents() {
		return 0, ErrMoneyUnderflow
	}
	return Money(rightSide.Cents() + money.Cents()), nil
}
