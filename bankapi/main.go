package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	bank "github.com/shubham-rewale/online-bank-project/bankcore"
)

var accounts map[float64]*bank.Account

func statementHandleFunc(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")

	if numberqs == "" {
		fmt.Fprintf(w, "Account Number is Missing")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid Account Number")
		return
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account wiht number %v can not be found", number)
		} else {
			fmt.Fprintf(w, "%s", account.Statement())
		}
	}
}

func depositHandleFunc(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	amountqs := r.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account Number is Missing")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid Account Number")
		return
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid Amount")
		return
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account wiht number %v can not be found", number)
			return
		} else {
			err := account.Deposit(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			} else {
				fmt.Fprintf(w, "%s", account.Statement())
			}
		}
	}
}

func withdrawHandleFunc(w http.ResponseWriter, r *http.Request) {
	numberqs := r.URL.Query().Get("number")
	amountqs := r.URL.Query().Get("amount")

	if numberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if number, err := strconv.ParseFloat(numberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid account number!")
		return
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
		return
	} else {
		account, ok := accounts[number]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", number)
			return
		} else {
			err := account.Withdraw(amount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			} else {
				fmt.Fprintf(w, "%s", account.Statement())
			}
		}
	}
}

func transferHandleFunc(w http.ResponseWriter, r *http.Request) {
	fromNumberqs := r.URL.Query().Get("fromNumber")
	toNumberqs := r.URL.Query().Get("toNumber")
	amountqs := r.URL.Query().Get("amount")

	if fromNumberqs == "" || toNumberqs == "" {
		fmt.Fprintf(w, "Account number is missing!")
		return
	}

	if fromNumber, err := strconv.ParseFloat(fromNumberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid source account number!")
		return
	} else if toNumber, err := strconv.ParseFloat(toNumberqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid destination account number!")
		return
	} else if amount, err := strconv.ParseFloat(amountqs, 64); err != nil {
		fmt.Fprintf(w, "Invalid amount number!")
		return
	} else {
		fromAccount, ok := accounts[fromNumber]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", fromNumber)
			return
		}

		toAccount, ok := accounts[toNumber]
		if !ok {
			fmt.Fprintf(w, "Account with number %v can't be found!", toNumber)
			return
		} else {
			err := fromAccount.TransferFunds(amount, toAccount)
			if err != nil {
				fmt.Fprintf(w, "%v", err)
				return
			} else {
				fmt.Fprintf(w, "%s", fromAccount.Statement())
				fmt.Fprintf(w, "%s", toAccount.Statement())
			}
		}

	}

}

func main() {
	accounts = map[float64]*bank.Account{
		1001: {
			Customer: bank.Customer{
				Name:    "John",
				Address: "Los Angeles, California",
				Phone:   "(213) 555 0147",
			},
			Number:  1001,
			Balance: 1000},
		1002: {
			Customer: bank.Customer{
				Name:    "Smith",
				Address: "Brooklyn, New York",
				Phone:   "(213) 255 0233",
			},
			Number:  1002,
			Balance: 0},
	}
	http.HandleFunc("/statement", statementHandleFunc)
	http.HandleFunc("/deposit", depositHandleFunc)
	http.HandleFunc("/withdraw", withdrawHandleFunc)
	http.HandleFunc("/transfer", transferHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
