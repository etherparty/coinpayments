package coinpayments

// Config is used by Viper to parse our provided credentials
type Config struct {
	PrivateKey           string `mapstructure:"private_key" json:"private_key"`
	PublicKey            string `mapstructure:"public_key" json:"public_key"`
	IPNSecret            string `mapstructure:"ipn_secret" json:"ipn_secret"`
	IPNURL               string `mapstructure:"ipn_url" json:"ipn_url"`
	BTCForwardingAddress string `json:"btc_forwarding_address,omitempty"`
	ETHForwardingAddress string `json:"eth_forwarding_address,omitempty"`
}
