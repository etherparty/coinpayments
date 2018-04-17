package coinpayments

import (
	"io"
	"io/ioutil"
	"net/url"
)

// IPNDepositResponse is a representation of a response received when any update is happening on a deposit
type IPNDepositResponse struct {
	Address    string `json:"address"`
	TxnID      string `json:"txn_id"`
	Status     string `json:"status"`
	StatusText string `json:"status_text"`
	Currency   string `json:"currency"`
	Confirms   string `json:"confirms"`
	Amount     string `json:"amount"`
	AmountI    string `json:"amounti"` // amount in satoshis
	Fee        string `json:"fee"`
	FeeI       string `json:"feei"` // fee in satoshis
	DestTag    string `json:"dest_tag"`
}

// HandleIPNDeposit takes a http request and returns the post parameters. io.Reader is typically fulfilled by the `req.Body` if used in a HTTP request to handle
// the IPN methods.
// IE: cps.HandleIPNDeposit(req.Body)
func (c *Client) HandleIPNDeposit(reader io.Reader) (*IPNDepositResponse, error) {

	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	return &IPNDepositResponse{
		Address:    values.Get("address"),
		TxnID:      values.Get("txn_id"),
		Status:     values.Get("status"),
		StatusText: values.Get("status_text"),
		Currency:   values.Get("currency"),
		Confirms:   values.Get("confirms"),
		Amount:     values.Get("amount"),
		AmountI:    values.Get("amounti"),
		Fee:        values.Get("fee"),
		FeeI:       values.Get("feei"),
		DestTag:    values.Get("dest_tag"),
	}, nil
}
