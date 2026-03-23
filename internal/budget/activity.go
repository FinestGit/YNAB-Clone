package budget

import "time"

func CategoryActivityForMonth(transactions []Transaction, categoryID CategoryID, month string) (Money, error) {
	var monthTotal Money = Money(0)
	for _, t := range transactions {
		if t.CategoryID == nil {
			continue
		}
		if *t.CategoryID != categoryID {
			continue
		}
		if monthKey(t.Date) != month {
			continue
		}
		nextTotal, err := monthTotal.Add(t.AmountCents)
		if err != nil {
			return 0, err
		}
		monthTotal = nextTotal
	}
	return monthTotal, nil
}

func monthKey(t time.Time) string {
	return t.Format("2006-01")
}
