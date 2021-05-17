//Procedural paradigm
package main

import (
	"errors"
	"fmt"
	"strconv"
)

type account struct {
	accountName   string
	accountNumber int
	balance       int
}

//Global variables
//Procedures communicate via global variables
var accounts []*account
var accountIdx int

func addAccount(name string) int {
	accountIdx++ //So that every account number is unique
	initialBalance := 0

	//Create new account
	newAcc := account{
		accountName:   name,
		accountNumber: accountIdx,
		balance:       initialBalance,
	}
	accounts = append(accounts, &newAcc)
	return newAcc.accountNumber
}

func getAccount(accountNum int) (*account, error) {
	for _, ac := range accounts {
		if ac.accountNumber == accountNum {
			return ac, nil
		}
	}
	return nil, errors.New("Account " + strconv.Itoa(accountNum) + " not found")
}

func deposit(accountNum, amt int) error {
	if ac, err := getAccount(accountNum); err == nil {
		ac.balance = ac.balance + amt
		return nil
	} else {
		return err
	}
}

func withdraw(accountNum, amt int) error {
	if ac, err := getAccount(accountNum); err == nil {
		ac.balance = ac.balance - amt
		return nil
	} else {
		return err
	}
}

func transfer(payerAcNum, beneficiaryAcNum, amt int) error {
	//Checking whether payer account exists
	if _, err := getAccount(payerAcNum); err == nil {
		//Checking whether benficiary account exists
		if _, err := getAccount(beneficiaryAcNum); err == nil {
			withdraw(payerAcNum, amt)
			deposit(beneficiaryAcNum, amt)
			return nil
		} else {
			return err
		}
	} else {
		return err
	}
}

func intro() {
	fmt.Printf("\nWelcome to monkey bank..." +
		"\nPress 1 to add new account" +
		"\nPress 2 to deposit money to account" +
		"\nPress 3 to withdraw money from account" +
		"\nPress 4 to transfer money from one account to another" +
		"\nPress 5 for balance enquiry" +
		"\nPress e to exit\n")
}

func main() {
	//Infinite loop
	for {
		intro()
		var userInput string
		fmt.Scanln(&userInput)

		switch userInput {
		case "1":
			var accName string
			fmt.Printf("\nPlease enter new account holder's name:")
			fmt.Scanln(&accName)
			acNum := addAccount(accName)
			fmt.Printf("\nNew account %d added; total accounts available are: %d", acNum, len(accounts))

		case "2":
			var accNum int
			fmt.Printf("\nPlease enter account holder's account number:")
			fmt.Scanln(&accNum)

			var amt int
			fmt.Printf("\nPlease enter Amount to deposit in Rs. :")
			fmt.Scanln(&amt)

			if err := deposit(accNum, amt); err != nil {
				fmt.Printf("\nError %s in adding; Rs. %d added to %d account",
					err.Error(), amt, accNum)
			} else {
				fmt.Printf("\nRs. %d added to %d account", amt, accNum)
			}
		case "3":
			var accNum int
			fmt.Printf("\nPlease enter account holder's account number:")
			fmt.Scanln(&accNum)

			var amt int
			fmt.Printf("\nPlease enter Amount to withdraw in Rs. :")
			fmt.Scanln(&amt)

			if err := withdraw(accNum, amt); err != nil {
				fmt.Printf("\nError %s in withdrawing; Rs. %d withdrawal from %d account",
					err.Error(), amt, accNum)
			} else {
				fmt.Printf("\nRs. %d withdrawal from %d account", amt, accNum)
			}
		case "4":
			var accNum int
			fmt.Printf("\nPlease enter payer account holder's account number:")
			fmt.Scanln(&accNum)

			var bNum int
			fmt.Printf("\nPlease enter beneficiary account holder's account number:")
			fmt.Scanln(&bNum)

			var amt int
			fmt.Printf("\nPlease enter Amount to transfer in Rs. :")
			fmt.Scanln(&amt)

			if err := transfer(accNum, bNum, amt); err != nil {
				fmt.Printf("\nError %s in transfering money; Rs. %d from %d account to %d account",
					err.Error(), amt, accNum, bNum)
			} else {
				fmt.Printf("\nRs. %d withdrawal from %d account", amt, accNum)
			}
		case "5":
			var accNum int
			fmt.Printf("\nPlease enter account holder's account number:")
			fmt.Scanln(&accNum)

			if ac, err := getAccount(accNum); err != nil {
				fmt.Printf("\nError %s in getting %d account",
					err.Error(), accNum)
			} else {
				fmt.Printf("\nBalance of %s(%d) account is %d",
					ac.accountName, accNum, ac.balance)
			}
		case "e":
			break
		}
	}

}
