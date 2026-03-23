package budget

type BudgetLine struct {
	Month      string
	CategoryID CategoryID
	Budgeted   Money
}

func (line BudgetLine) Available(activity Money) (Money, error) {
	available, err := line.Budgeted.Add(activity)
	if err != nil {
		return 0, err
	}
	return available, nil
}
