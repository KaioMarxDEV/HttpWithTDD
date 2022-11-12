package bankcore

import (
	"encoding/json"
	"testing"
)

func TestAccount(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if account.Name == "" {
		t.Error("can't create an Account object")
	}
}

func TestDeposit(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if account.Balance != 10 {
		t.Error("balance is not being updated after a deposit")
	}
}

func TestDepositError(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	if err := account.Deposit(-1); err == nil {
		t.Error("only positive numbers should be allowed to deposit")
	}
}

func TestWithdraw(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)
	account.Withdraw(9)

	if account.Balance != 1 {
		t.Error("Failed in withdraw values from account")
	}
}

func TestWithdrawError(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	account.Deposit(10)

	if err := account.Withdraw(11); err == nil {
		t.Error("should not be able to withdraw values greater than avaiable")
	}
}

func TestStatement(t *testing.T) {
	account := Account{
		Customer: Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 0,
	}

	statement, _ := account.Statement()

	if !json.Valid(statement) {
		t.Error("statement format is invalid")
	}
}
