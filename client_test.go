package coinpayments_test

import (
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/etherparty/coinpayments"
)

func testClient() (*coinpayments.Client, error) {
	return coinpayments.NewClient(&coinpayments.Config{PublicKey: os.Getenv("COINPAYMENTS_PUBLIC_KEY"), PrivateKey: os.Getenv("COINPAYMENTS_PRIVATE_KEY")}, &http.Client{})

}
func TestNewClient(t *testing.T) {
	if _, err := coinpayments.NewClient(&coinpayments.Config{PublicKey: "", PrivateKey: ""}, &http.Client{}); err == nil {
		t.Fatalf("Should have thrown an error with emptu public and private key, but it didn't")
	}

	if _, err := coinpayments.NewClient(&coinpayments.Config{PublicKey: "publickey", PrivateKey: "privateKey"}, nil); err == nil {
		t.Fatalf("Should have thrown an error with a nil http client, but it didn't")
	}

	if _, err := testClient(); err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}
}

func TestCall(t *testing.T) {
	client, err := testClient()
	if err != nil {
		t.Fatalf("Should have instantiated a new client with valid config and http client, but it threw error: %s", err.Error())
	}

	failCommand := "doesntexist"
	var txResponse coinpayments.TransactionResponse
	if err := client.Call(failCommand, url.Values{}, &txResponse); err == nil {
		t.Fatalf("Should have failed with non-supported command of %s", failCommand)
	}

}
