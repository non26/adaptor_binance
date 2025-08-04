package serviceconfig

type ServiceConfig struct {
	Secret Secrets `mapstructure:"secret" json:"secret"`
	Future Future  `mapstructure:"future" json:"future"`
	Spot   Spot    `mapstructure:"spot" json:"spot"`
}

type Secrets struct {
	Account1 AccountSecret `mapstructure:"account1" json:"account1"`
	Account2 AccountSecret `mapstructure:"account2" json:"account2"`
	Account3 AccountSecret `mapstructure:"account3" json:"account3"`
}

type AccountSecret struct {
	ApiKey    string `mapstructure:"apiKey" json:"apiKey"`
	ApiSecret string `mapstructure:"apiSecret" json:"apiSecret"`
}

type Future struct {
	Usdm  FutureDetail `mapstructure:"usdm" json:"usdm"`
	Coinm FutureDetail `mapstructure:"coinm" json:"coinm"`
}

type FutureDetail struct {
	Url                        string `mapstructure:"url" json:"url"`
	NewOrderEndpoint           string `mapstructure:"newOrderEndpoint" json:"newOrderEndpoint"`
	PlaceMultipleOrderEndpoint string `mapstructure:"placeMultipleOrderEndpoint" json:"placeMultipleOrderEndpoint"`
}

type Spot struct {
	Url                        string `mapstructure:"url" json:"url"`
	NewOrderEndpoint           string `mapstructure:"newOrderEndpoint" json:"newOrderEndpoint"`
	PlaceMultipleOrderEndpoint string `mapstructure:"placeMultipleOrderEndpoint" json:"placeMultipleOrderEndpoint"`
}
