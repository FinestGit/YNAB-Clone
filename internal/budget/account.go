package budget

import (
	"errors"
	"strings"
)

type Account struct {
	ID      AccountID
	Name    string
	Balance Money
}

var ErrEmptyAccountName = errors.New("account name cannot be empty")

func NewAccount(id AccountID, name string, balance Money) (Account, error) {
	var accountName string = strings.TrimSpace(name)
	if accountName == "" {
		return Account{}, ErrEmptyAccountName
	}

	return Account{
		ID:      AccountID(id),
		Name:    accountName,
		Balance: Money(balance),
	}, nil
}
