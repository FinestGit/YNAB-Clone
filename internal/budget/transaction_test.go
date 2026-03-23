package budget

import (
	"errors"
	"math"
	"testing"
	"time"
)

func TestApplyTransactionToAccount_Inflow(t *testing.T) {
	accountID, _ := NewAccountID("test")
	var accountName string = "test"
	account, _ := NewAccount(accountID, accountName, Money(100))
	var newTransaction Transaction = Transaction{
		ID:          "test",
		AccountID:   accountID,
		AmountCents: Money(50),
		Date:        time.Now(),
		payee:       "test",
	}
	newAccount, err := ApplyTransactionToAccount(account, newTransaction)
	if err != nil {
		t.Fatal(err)
	}
	if newAccount.ID.String() != accountID.String() {
		t.Fatalf("new account id is %s, should be %s", newAccount.ID.String(), accountID.String())
	}
	if newAccount.Name != accountName {
		t.Fatalf("new account name is %s, should be %s", newAccount.Name, accountName)
	}
	if newAccount.Balance != 150 {
		t.Fatalf("new account balance is %d, should be 150", newAccount.Balance)
	}
}

func TestApplyTransactionToAccount_Outflow(t *testing.T) {
	accountID, _ := NewAccountID("test")
	var accountName string = "test"
	account, _ := NewAccount(accountID, accountName, Money(100))
	var newTransaction Transaction = Transaction{
		ID:          "test",
		AccountID:   accountID,
		AmountCents: Money(-50),
		Date:        time.Now(),
		payee:       "test",
	}
	newAccount, err := ApplyTransactionToAccount(account, newTransaction)
	if err != nil {
		t.Fatal(err)
	}
	if newAccount.Balance != 50 {
		t.Fatalf("new account balance is %d, should be 50", newAccount.Balance)
	}
}

func TestApplyTransactionToAccount_ZeroError(t *testing.T) {
	accountID, _ := NewAccountID("test")
	var accountName string = "test"
	account, _ := NewAccount(accountID, accountName, Money(100))
	var newTransaction Transaction = Transaction{
		ID:          "test",
		AccountID:   accountID,
		AmountCents: Money(0),
		Date:        time.Now(),
		payee:       "test",
	}
	_, err := ApplyTransactionToAccount(account, newTransaction)
	if !errors.Is(err, ErrZeroTransactionAmount) {
		t.Fatal("ApplyTransactionToAccount should return ErrZeroTransactionAmount when AmountCents is 0")
	}
}

func TestApplyTransactionToAccount_AccountMismatchError(t *testing.T) {
	accountID, _ := NewAccountID("test")
	var accountName string = "test"
	account, _ := NewAccount(accountID, accountName, Money(100))
	var newTransaction Transaction = Transaction{
		ID:          "test",
		AccountID:   "other_test",
		AmountCents: Money(-50),
		Date:        time.Now(),
		payee:       "test",
	}
	_, err := ApplyTransactionToAccount(account, newTransaction)
	if !errors.Is(err, ErrTransactionAccountMismatch) {
		t.Fatal("ApplyTransactionToAccount should return ErrTransactionAccountMismatch when AccountIDs are not the same")
	}
}

func TestApplyTransactionToAccount_BubbleUpError(t *testing.T) {
	accountID, _ := NewAccountID("test")
	var accountName string = "test"
	account, _ := NewAccount(accountID, accountName, Money(math.MaxInt64))
	var newTransaction Transaction = Transaction{
		ID:          "test",
		AccountID:   accountID,
		AmountCents: Money(1),
		Date:        time.Now(),
		payee:       "test",
	}
	_, err := ApplyTransactionToAccount(account, newTransaction)
	if !errors.Is(err, ErrMoneyOverflow) {
		t.Fatal("ApplyTransactionToAccount should return ErrMoneyOverflow when adding to a max int account balance")
	}
}
