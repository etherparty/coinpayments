package coinpayments_test

import (
	"fmt"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/etherparty/coinpayments"
)

func TestCallGetConversionLimits(t *testing.T) {
	pub, ok := os.LookupEnv("COINPAYMENTS_PUBLIC_KEY")
	if !ok {
		t.Fatal("no public key provided in environment")
	}

	priv, ok := os.LookupEnv("COINPAYMENTS_PRIVATE_KEY")
	if !ok {
		t.Fatal("no priatekey provided in environment")
	}

	cfg := &coinpayments.Config{
		PrivateKey: priv,
		PublicKey:  pub,
	}
	hc := &http.Client{Timeout: time.Minute * 1}

	client, err := coinpayments.NewClient(cfg, hc)
	if err != nil {
		t.Fatal(err)
	}

	req1 := coinpayments.ConvertLimitRequest{
		From: "ETH",
		To:   "BTC",
	}
	req2 := coinpayments.ConvertLimitRequest{
		From: "BTC",
		To:   "ETH",
	}

	resp1, err := client.CallGetConversionLimits(&req1)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Printf("Conversion limits for ETH->FUEL\n%+v\n", resp1)
	}

	resp2, err := client.CallGetConversionLimits(&req2)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Conversion limits for BTC->FUEL\n%+v\n", resp2)

}
