package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	bank "github.com/test/bankcore"
)

var accounts = map[float64]*bank.Account{}

func main() {
	accounts[1001] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number:  1001,
		Balance: 150,
	}
	accounts[2002] = &bank.Account{
		Customer: bank.Customer{
			Name:    "John",
			Address: "Los Angeles, California",
			Phone:   "(213) 555 0147",
		},
		Number: 2002,
	}

	http.HandleFunc("/statement", statement)
	http.HandleFunc("/deposit", deposit)
	http.HandleFunc("/withdraw", withdraw)
	http.HandleFunc("/transfer", transfer)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func statement(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			json, _ := account.Statement()
			fmt.Fprintf(w, "Response: %v", json)
		}
	}
}

func deposit(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				json, _ := account.Statement()
				fmt.Fprintf(w, "Data: %v", json)
			}
		}
	}
}

func withdraw(w http.ResponseWriter, req *http.Request) {
	numberqs := req.URL.Query().Get("number")
	amountqs := req.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
			} else {
				json, _ := account.Statement()
				fmt.Fprintf(w, "Response: %v", json)
			}
		}
	}
}

func transfer(w http.ResponseWriter, req *http.Request) {
	senderqs := req.URL.Query().Get("sender")
	receiverqs := req.URL.Query().Get("receiver")
	amountqs := req.URL.Query().Get("amount")

	if senderqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if receiverqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if sender, err := strconv.ParseFloat(senderqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid sender account number!")
	} else if receiver, err := strconv.ParseFloat(receiverqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid receiver account number!")
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
	} else {
		sendAccount, ok := accounts[sender]

		if !ok {
			fmt.Fprintf(w, "Your account was not found!")
		} else {
			receiverAccount, ok := accounts[receiver]

			if !ok {
				fmt.Fprintf(w, "Receiver account was not found!")
			} else {
				senderErr := sendAccount.Withdraw(amount)

				if senderErr != nil {
					fmt.Fprintf(w, "Failed to transfer values, insuficient values to withdraw")
				} else {
					receiverErr := receiverAccount.Deposit(amount)
					if receiverErr != nil {
						fmt.Fprintf(w, "Failed to transfer values, receiver cannot accept deposits")
					} else {
						senderJson, _ := sendAccount.Statement()
						receiverJson, _ := receiverAccount.Statement()
						fmt.Fprintf(w, "Sender: %v ### Receiver: %v", senderJson, receiverJson)
					}
				}
			}
		}
	}
}
