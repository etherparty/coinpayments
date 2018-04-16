package coinpayments_test

import (
	"testing"

	"github.com/etherparty/coinpayments"
)

func TestHandleIPNDeposit(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	resp, err := client.HandleIPNDeposit(&coinpayments.Reader{Data: []byte("{\"status\":\"200\"}")})

	if err != nil {
		t.Fatalf("Could not handle IPN deposit with mock reader: got %s", err.Error())
	}

	if resp.Status != "200" {
		t.Fatalf("Expected status %s, got %s", "200", resp.Status)
	}
}
