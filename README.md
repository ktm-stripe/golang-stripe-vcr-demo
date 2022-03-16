# Stripe VCR Demo in Golang

## Instructions

This is meant to be run with Docker. You'll need a secret key and your `account_id`, findable at https://dashboard.stripe.com/test/apikeys and https://dashboard.stripe.com/settings/account respectively.

```shell
$ docker build -t golang-stripe-vcr-demo
$ docker run -e STRIPE_KEY=sk_test_xxxx -e STRIPE_ACCOUNT=acct_xxxx golang-stripe-vcr-demo
```

Run the tests with

```shell
$ docker run -e STRIPE_KEY=sk_test_xxxx -e STRIPE_ACCOUNT=acct_xxxx golang-stripe-vcr-demo go test -v
```

If working correctly, you should be able to disconnect from the network and tests should still pass.
