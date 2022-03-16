package main

import (
	"fmt"
	"os"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/account"
)

func main() {
	stripe.Key = os.Getenv("STRIPE_KEY")

	a, _ := GetAccount()

	fmt.Println(a)
}

func Msg() string {
	return "Success"
}

func GetAccount() (*stripe.Account, error) {
	a, err := account.GetByID(
		os.Getenv("STRIPE_ACCOUNT"),
		nil,
	)

	return a, err
}
