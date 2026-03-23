package budget

import "errors"

type BudgetLine struct {
	Month      string
	CategoryID CategoryID
	Budgeted   Money
}

var ErrInvalidAllocateAmount = errors.New("cannot allocate 0 or negative amounts")

func (line BudgetLine) Available(activity Money) (Money, error) {
	available, err := line.Budgeted.Add(activity)
	if err != nil {
		return 0, err
	}
	return available, nil
}

func (line BudgetLine) Allocate(amount Money) (BudgetLine, error) {
	if amount.Cents() <= 0 {
		return BudgetLine{}, ErrInvalidAllocateAmount
	}
	newBudgeted, err := line.Budgeted.Add(amount)
	if err != nil {
		return BudgetLine{}, err
	}
	return BudgetLine{
		Month:      line.Month,
		CategoryID: line.CategoryID,
		Budgeted:   newBudgeted,
	}, nil
}
