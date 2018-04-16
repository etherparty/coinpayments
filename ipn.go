package coinpayments

import (
	"io"
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

	var resp IPNDepositResponse
	if err := c.unmarshal(reader, &resp); err != nil {
		return nil, err
	}

	return &resp, nil
}
