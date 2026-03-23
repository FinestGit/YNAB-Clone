package budget

import (
	"errors"
	"math"
	"testing"
	"time"
)

func TestAvailable_BudgetLineSpending(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	var budgetLine BudgetLine = BudgetLine{
		Month:      time.Now().Format("2006-01"),
		CategoryID: categoryID,
		Budgeted:   Money(100),
	}
	available, err := budgetLine.Available(Money(-30))
	if err != nil {
		t.Fatal(err)
	}
	if available != 70 {
		t.Fatalf("available is %d, should be 70", available)
	}
}

func TestAvailable_ZeroActivity(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	var budgetLine BudgetLine = BudgetLine{
		Month:      time.Now().Format("2006-01"),
		CategoryID: categoryID,
		Budgeted:   Money(100),
	}
	available, err := budgetLine.Available(Money(0))
	if err != nil {
		t.Fatal(err)
	}
	if available != 100 {
		t.Fatalf("available is %d, should be 100", available)
	}
}

func TestAvailable_BubbleUpError(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	var budgetLine BudgetLine = BudgetLine{
		Month:      time.Now().Format("2006-01"),
		CategoryID: categoryID,
		Budgeted:   Money(math.MaxInt64),
	}
	_, err := budgetLine.Available(Money(1))
	if !errors.Is(err, ErrMoneyOverflow) {
		t.Fatal("Available should throw ErrMoneyOverflow when budgeted is maxInt64 plus activity with 1")
	}
}
