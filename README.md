# coinpayments

Coinpayments Golang wrapper. To use, instantiate a `*coinpayments.Client` by calling the `coinpayments.NewClient` function. You will need to pass in a valid instance of
the `*coinpayments.Config` struct, consisting of your private and public keys, as well as an instance of your desired http client.  
  
This looks like: 
`client, err := coinpayments.NewClient(&coinpayments.Config{ PublicKey: "yourpublickey", PrivateKey: "yourprivatekey"}, &http.Client{ Timeout: 10 * time.Second})`.  
  
Once instantiated, you can use the client to make API calls. An example of the `create_transaction` command being called:  
  
  `client.CallCreateTransaction(amount, currency1, currency2, buyerEmail)`, where the amount is the value of the fiat currency you wish to receive.  
  currency1 is the type of currency that amount is specified in (USD, CAD, JPY, etc).  
  currency2 is the type of currency in crypto that you wish to receive (BTC, ETC, LTC, etc).  
  buyerEmail is optional with the API but mandatory for our client.  

# Calling commands on the API
If you don't have a valid Reader to pass in to the various Call methods like a `req.Body`,
you can create an instance of the Reader and pass in your data as a byte slice.
ie: 
```
client.CallGetDepositAddress(&coinpayments.Reader{data: []byte("{\"currency\":\"USD\"}")}`  
```  

# Call
If you don't want to pass a Reader to the individual functions, you can use Call to call any command directly.
You just need to build out your own data with an url.Values, and create an instance of the responsestruct expected by your command.
ie: 
```
data := url.Values{}
data.Add("currency", "USD")
resp := coinpayments.DepositAddressResponse{}
err := c.Call(coinpayments.CmdGetDepositAddress, data, &resp)
```

# tests

You need to export two environment variables for the tests to run - your public key, and private key.  
```
export COINPAYMENTS_PUBLIC_KEY=yourpublickeyhere
export COINPAYMENTS_PRIVATE_KEY=yourprivatekeyhere  

go test ./...
```
