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
	priv := os.Getenv("PRIVATE_KEY")
	if priv == "" {
		t.Fatal("env PRIVATE_KEY is not set")
	}
	pub := os.Getenv("PUBLIC_KEY")
	if pub == "" {
		t.Fatal("env PUBLIC_KEY is not set")
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
		To:   "FUEL",
	}
	req2 := coinpayments.ConvertLimitRequest{
		From: "BTC",
		To:   "FUEL",
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
