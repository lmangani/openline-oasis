package config

type Config struct {
	Service struct {
		MessageStore  string `env:"MESSAGE_STORE_URL,required"`
		ServerAddress string `env:"CHANNELS_API_SERVER_ADDRESS,required"`
		GRPCPort      int    `env:"CHANNELS_GRPC_PORT,required"`
		OasisApiUrl   string `env:"OASIS_API_URL,required"`
	}
	Mail struct {
		SMTPSeverAddress  string `env:"SMTP_SERVER_ADDRESS,required"`
		SMTPSeverUser     string `env:"SMTP_SERVER_USER,required"`
		SMTPSeverPassword string `env:"SMTP_SERVER_PASSWORD,required"`
		SMTPFromUser      string `env:"SMTP_FROM_USER,required"`
		ApiKey            string `env:"MAIL_API_KEY,required"`
	}
}
