package budget

import (
	"errors"
	"time"
)

type Transaction struct {
	ID          TransactionID
	AccountID   AccountID
	CategoryID  *CategoryID
	AmountCents Money
	Date        time.Time
	payee       string
}

var ErrZeroTransactionAmount = errors.New("cannot have a transaction amount of 0")
var ErrTransactionAccountMismatch = errors.New("transaction account id must be equal to passed in account")

func ApplyTransactionToAccount(account Account, transaction Transaction) (Account, error) {
	if transaction.AmountCents == 0 {
		return Account{}, ErrZeroTransactionAmount
	}
	if account.ID.String() != transaction.AccountID.String() {
		return Account{}, ErrTransactionAccountMismatch
	}
	newBalance, err := account.Balance.Add(transaction.AmountCents)
	if err != nil {
		return Account{}, err
	}
	return Account{
		ID:      AccountID(account.ID),
		Name:    account.Name,
		Balance: Money(newBalance),
	}, nil
}
