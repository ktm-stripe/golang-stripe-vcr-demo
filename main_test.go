package main

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/dnaeon/go-vcr/cassette"
	"github.com/dnaeon/go-vcr/recorder"
	"github.com/stripe/stripe-go"
)

func TestMsg(t *testing.T) {
	out := Msg()
	if out != "Success" {
		t.Errorf("Msg() = \"%s\", expected \"Success\"", out)
	}
}

func stripeVCR(tb testing.TB) func(tb testing.TB) {
	r, _ := recorder.New("fixtures/stripe")
	c := &http.Client{Transport: r}
	stripe.SetHTTPClient(c)
	stripe.Key = os.Getenv("STRIPE_KEY")

	// Add a filter which removes Authorization headers from all requests:
	r.AddFilter(func(i *cassette.Interaction) error {
		delete(i.Request.Headers, "Authorization")
		return nil
	})

	return func(tb testing.TB) {
		r.Stop()
	}
}

func TestGetAccount(t *testing.T) {
	teardown := stripeVCR(t)
	defer teardown(t)

	a, err := GetAccount()
	fmt.Println(a)
	if err != nil {
		t.Error()
	}
}
