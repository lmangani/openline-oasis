package config

type Config struct {
	Service struct {
		MessageStore string `env:"MESSAGE_STORE_URL,required"`
	}
}