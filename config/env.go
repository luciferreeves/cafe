package config

type server struct {
	Host      string `env:"SERVER_HOST" default:"localhost"`
	Port      int    `env:"SERVER_PORT" default:"8080"`
	AppSecret string `env:"APP_SECRET" default:"mysecret"`
	DevMode   bool   `env:"DEV_MODE" default:"true"`
}
