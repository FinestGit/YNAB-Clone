package budget

import (
	"errors"
	"testing"
)

func TestNewAccount_Success(t *testing.T) {
	accountID, _ := NewAccountID("test")
	balance, _ := NewMoney(100)
	var accountName string = "test"
	newAccount, err := NewAccount(accountID, accountName, balance)
	if err != nil {
		t.Fatal(err)
	}
	if newAccount.Balance.Cents() != balance.Cents() {
		t.Fatalf("balance is %d, should be %d", newAccount.Balance.Cents(), balance.Cents())
	}
	if newAccount.ID.String() != accountID.String() {
		t.Fatalf("account id is %s, should be %s", newAccount.ID.String(), accountID.String())
	}
	if newAccount.Name != accountName {
		t.Fatalf("account name is %s, should be %s", newAccount.Name, accountName)
	}
}

func TestNewAccount_EmptyName(t *testing.T) {
	accountID, _ := NewAccountID("test")
	balance, _ := NewMoney(100)
	newAccount, err := NewAccount(accountID, " ", balance)
	if !errors.Is(err, ErrEmptyAccountName) {
		t.Fatalf("Account name is %s, should have received ErrEmptyAccountName", newAccount.Name)
	}
}
