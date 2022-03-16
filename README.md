# Stripe VCR Demo in Golang

## Instructions

This is meant to be run with Docker, and you'll need a secret key from https://dashboard.stripe.com/test/apikeys.

```shell
$ docker build -t golang-stripe-vcr-demo .
$ docker run -e STRIPE_KEY=sk_test_xxxx golang-stripe-vcr-demo
```

Run the tests with

```shell
$ docker run -e STRIPE_KEY=sk_test_xxxx golang-stripe-vcr-demo go test -v
```

If working correctly, you should be able to disconnect from the network and tests should still pass.
