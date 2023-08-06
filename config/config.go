package config

var Cfg Config

type Config struct {
	Transaction Transaction `mapstructure:"transaction"`
}

type Transaction struct {
	MaxDescriptionLength int `mapstructure:"max_description_length"`
}

func init() {
	Cfg = Config{
		Transaction: Transaction{
			MaxDescriptionLength: 100,
		},
	}
}
