package budget

import (
	"errors"
	"testing"
)

func TestNewAccountID_RoundTrip(t *testing.T) {
	var testString string = "test"
	accountID, err := NewAccountID(testString)
	if err != nil {
		t.Fatal(err)
	}
	if accountID.String() != testString {
		t.Fatalf("%s != %s", accountID.String(), testString)
	}
}

func TestNewAccountID_EmptyString(t *testing.T) {
	_, err := NewAccountID("")
	if !errors.Is(err, ErrEmptyAccountID) {
		t.Fatal("expected ErrEmptyAccountID")
	}
}

func TestNewCategoryID_RoundTrip(t *testing.T) {
	var testString string = "test"
	categoryID, err := NewCategoryID(testString)
	if err != nil {
		t.Fatal(err)
	}
	if categoryID.String() != testString {
		t.Fatalf("%s != %s", categoryID.String(), testString)
	}
}

func TestNewCategoryID_EmptyString(t *testing.T) {
	_, err := NewCategoryID("")
	if !errors.Is(err, ErrEmptyCategoryID) {
		t.Fatal("expected ErrEmptyCategoryID")
	}
}

func TestNewTransactionID_RoundTrip(t *testing.T) {
	var testString string = "test"
	transactionID, err := NewTransactionID(testString)
	if err != nil {
		t.Fatal(err)
	}
	if transactionID.String() != testString {
		t.Fatalf("%s != %s", transactionID.String(), testString)
	}
}

func TestNewTransactionID_EmptyString(t *testing.T) {
	_, err := NewTransactionID("")
	if !errors.Is(err, ErrEmptyTransactionID) {
		t.Fatal("expected ErrEmptyTransactionID")
	}
}
