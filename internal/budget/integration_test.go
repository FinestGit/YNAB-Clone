package budget

import (
	"testing"
	"time"
)

func TestBudgetGroceriesMatch_EndToEndActivityAndAvailable(t *testing.T) {
	var month string = "2025-03"
	firstCategoryID, _ := NewCategoryID("groceries")
	secondCategoryID, _ := NewCategoryID("gas")
	accountID, _ := NewAccountID("test-account")
	transactionOne := Transaction{
		ID:          "1",
		AccountID:   accountID,
		CategoryID:  &firstCategoryID,
		AmountCents: Money(-2_000),
		Date:        time.Date(2025, 3, 10, 0, 0, 0, 0, time.UTC),
		payee:       "test-payee",
	}
	transactionTwo := Transaction{
		ID:          "2",
		AccountID:   accountID,
		CategoryID:  &firstCategoryID,
		AmountCents: Money(500),
		Date:        time.Date(2025, 3, 15, 0, 0, 0, 0, time.UTC),
		payee:       "test-payee",
	}
	transactionThree := Transaction{
		ID:          "3",
		AccountID:   accountID,
		CategoryID:  &secondCategoryID,
		AmountCents: Money(-1_000),
		Date:        time.Date(2025, 3, 20, 0, 0, 0, 0, time.UTC),
		payee:       "test-payee",
	}
	transactionFour := Transaction{
		ID:          "4",
		AccountID:   accountID,
		CategoryID:  &firstCategoryID,
		AmountCents: Money(-9_999),
		Date:        time.Date(2025, 2, 28, 0, 0, 0, 0, time.UTC),
		payee:       "test-payee",
	}
	transactionFive := Transaction{
		ID:          "5",
		AccountID:   accountID,
		AmountCents: Money(-5_000),
		Date:        time.Date(2025, 3, 1, 0, 0, 0, 0, time.UTC),
		payee:       "test-payee",
	}
	var transactions []Transaction = []Transaction{transactionOne, transactionTwo, transactionThree, transactionFour, transactionFive}
	activity, err := CategoryActivityForMonth(transactions, firstCategoryID, month)
	if err != nil {
		t.Fatal(err)
	}
	if activity.Cents() != -1_500 {
		t.Fatalf("activity is %d, should be -1_500", activity.Cents())
	}
	line := BudgetLine{
		Month:      month,
		CategoryID: firstCategoryID,
		Budgeted:   Money(10_000),
	}
	available, err := line.Available(activity)
	if err != nil {
		t.Fatal(err)
	}
	if available.Cents() != 8_500 {
		t.Fatalf("available is %d, should be 8_500", available.Cents())
	}
}
