package budget

import "errors"

type AccountID string
type CategoryID string
type TransactionID string

var ErrEmptyAccountID = errors.New("AccountID cannot be empty")
var ErrEmptyCategoryID = errors.New("CategoryID cannot be empty")
var ErrEmptyTransactionID = errors.New("TransactionID cannot be empty")

func NewAccountID(s string) (AccountID, error) {
	if s == "" {
		return "", ErrEmptyAccountID
	}
	return AccountID(s), nil
}

func (id AccountID) String() string {
	return string(id)
}

func NewCategoryID(s string) (CategoryID, error) {
	if s == "" {
		return "", ErrEmptyCategoryID
	}
	return CategoryID(s), nil
}

func (id CategoryID) String() string {
	return string(id)
}

func NewTransactionID(s string) (TransactionID, error) {
	if s == "" {
		return "", ErrEmptyTransactionID
	}
	return TransactionID(s), nil
}

func (id TransactionID) String() string {
	return string(id)
}
