package budget

import (
	"errors"
	"math"
	"testing"
	"time"
)

func TestCategoryActivityForMonth_MatchingCategoryAndMonth(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	accountID, _ := NewAccountID("test")
	var transactionOne Transaction = Transaction{
		ID:          TransactionID("test1"),
		AccountID:   accountID,
		CategoryID:  &categoryID,
		AmountCents: Money(50),
		Date:        time.Now(),
		payee:       "test",
	}
	var transactionTwo Transaction = Transaction{
		ID:          TransactionID("test2"),
		AccountID:   accountID,
		CategoryID:  &categoryID,
		AmountCents: Money(30),
		Date:        time.Now(),
		payee:       "test",
	}
	var month string = time.Now().Format("2006-01")
	var transactions []Transaction = []Transaction{transactionOne, transactionTwo}
	money, err := CategoryActivityForMonth(transactions, categoryID, month)
	if err != nil {
		t.Fatal(err)
	}
	if money.Cents() != 80 {
		t.Fatalf("value for category %s is %d, should be 80", categoryID.String(), money.Cents())
	}
}

func TestCategoryActivityForMonth_ExcludesDifferentCategory(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	secondCategoryID, _ := NewCategoryID("gas")
	accountID, _ := NewAccountID("test")
	var transactionOne Transaction = Transaction{
		ID:          TransactionID("test1"),
		AccountID:   accountID,
		CategoryID:  &categoryID,
		AmountCents: Money(50),
		Date:        time.Now(),
		payee:       "test",
	}
	var transactionTwo Transaction = Transaction{
		ID:          TransactionID("test2"),
		AccountID:   accountID,
		CategoryID:  &secondCategoryID,
		AmountCents: Money(30),
		Date:        time.Now(),
		payee:       "test",
	}
	var month string = time.Now().Format("2006-01")
	var transactions []Transaction = []Transaction{transactionOne, transactionTwo}
	money, err := CategoryActivityForMonth(transactions, categoryID, month)
	if err != nil {
		t.Fatal(err)
	}
	if money.Cents() != 50 {
		t.Fatalf("value for category %s is %d, should be 50", categoryID.String(), money.Cents())
	}
}

func TestCategoryActivityForMonth_ExcludeDifferentMonth(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	accountID, _ := NewAccountID("test")
	var transactionOne Transaction = Transaction{
		ID:          TransactionID("test1"),
		AccountID:   accountID,
		CategoryID:  &categoryID,
		AmountCents: Money(50),
		Date:        time.Now(),
		payee:       "test",
	}
	var transactionTwo Transaction = Transaction{
		ID:          TransactionID("test2"),
		AccountID:   accountID,
		CategoryID:  &categoryID,
		AmountCents: Money(30),
		Date:        time.Now().AddDate(0, -1, 0),
		payee:       "test",
	}
	var month string = time.Now().Format("2006-01")
	var transactions []Transaction = []Transaction{transactionOne, transactionTwo}
	money, err := CategoryActivityForMonth(transactions, categoryID, month)
	if err != nil {
		t.Fatal(err)
	}
	if money.Cents() != 50 {
		t.Fatalf("value for category %s is %d, should be 50", categoryID.String(), money.Cents())
	}
}

func TestCategoryActivityForMonth_ExcludeUncategorized(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	accountID, _ := NewAccountID("test")
	var transactionOne Transaction = Transaction{
		ID:          TransactionID("test1"),
		AccountID:   accountID,
		CategoryID:  &categoryID,
		AmountCents: Money(50),
		Date:        time.Now(),
		payee:       "test",
	}
	var transactionTwo Transaction = Transaction{
		ID:          TransactionID("test2"),
		AccountID:   accountID,
		AmountCents: Money(30),
		Date:        time.Now(),
		payee:       "test",
	}
	var month string = time.Now().Format("2006-01")
	var transactions []Transaction = []Transaction{transactionOne, transactionTwo}
	money, err := CategoryActivityForMonth(transactions, categoryID, month)
	if err != nil {
		t.Fatal(err)
	}
	if money.Cents() != 50 {
		t.Fatalf("value for category %s is %d, should be 50", categoryID.String(), money.Cents())
	}
}

func TestCategoryActivityForMonth_BubbleUpError(t *testing.T) {
	categoryID, _ := NewCategoryID("groceries")
	accountID, _ := NewAccountID("test")
	var transactionOne Transaction = Transaction{
		ID:          TransactionID("test1"),
		AccountID:   accountID,
		CategoryID:  &categoryID,
		AmountCents: Money(math.MaxInt64),
		Date:        time.Now(),
		payee:       "test",
	}
	var transactionTwo Transaction = Transaction{
		ID:          TransactionID("test2"),
		AccountID:   accountID,
		CategoryID:  &categoryID,
		AmountCents: Money(1),
		Date:        time.Now(),
		payee:       "test",
	}
	var month string = time.Now().Format("2006-01")
	var transactions []Transaction = []Transaction{transactionOne, transactionTwo}
	_, err := CategoryActivityForMonth(transactions, categoryID, month)
	if !errors.Is(err, ErrMoneyOverflow) {
		t.Fatal("CategoryActivityForMonth should return ErrMoneyOverflow when adding to a max int into category that has anything but 0")
	}
}
