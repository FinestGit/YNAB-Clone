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

func TestAllocate_BudgetLineAllocation(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	var budgetLine BudgetLine = BudgetLine{
		Month:      time.Now().Format("2006-01"),
		CategoryID: categoryID,
		Budgeted:   Money(100),
	}
	newBudgetLine, err := budgetLine.Allocate(Money(100))
	if err != nil {
		t.Fatal(err)
	}
	if newBudgetLine.Budgeted.Cents() != 200 {
		t.Fatalf("new budgeted after allocation is %d, should be 200", newBudgetLine.Budgeted.Cents())
	}
	if newBudgetLine.Month != budgetLine.Month {
		t.Fatalf("new budgeted month is %s, should be %s", newBudgetLine.Month, budgetLine.Month)
	}
	if newBudgetLine.CategoryID != budgetLine.CategoryID {
		t.Fatalf("new budgeted category id is %s, should be %s", newBudgetLine.CategoryID.String(), budgetLine.CategoryID.String())
	}
}

func TestAllocate_NegativeAllocationError(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	var budgetLine BudgetLine = BudgetLine{
		Month:      time.Now().Format("2006-01"),
		CategoryID: categoryID,
		Budgeted:   Money(100),
	}
	_, err := budgetLine.Allocate(Money(-1))
	if !errors.Is(err, ErrInvalidAllocateAmount) {
		t.Fatal("Allocate should throw ErrInvalidAllocateAmount when 0 or negative is provided")
	}
	_, err = budgetLine.Allocate(Money(0))
	if !errors.Is(err, ErrInvalidAllocateAmount) {
		t.Fatal("Allocate should throw ErrInvalidAllocateAmount when 0 or negative is provided")
	}
}

func TestAllocate_BubbleUpError(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	var budgetLine BudgetLine = BudgetLine{
		Month:      time.Now().Format("2006-01"),
		CategoryID: categoryID,
		Budgeted:   Money(math.MaxInt64),
	}
	_, err := budgetLine.Allocate(Money(1))
	if !errors.Is(err, ErrMoneyOverflow) {
		t.Fatal("Allocate should throw ErrMoneyOverflow when budgeted is maxInt64 and allocation occurs")
	}
}
