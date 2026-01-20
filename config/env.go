package config

type server struct {
	Host      string `env:"SERVER_HOST" default:"localhost"`
	Port      int    `env:"SERVER_PORT" default:"8080"`
	AppSecret string `env:"APP_SECRET" default:"mysecret"`
	DevMode   bool   `env:"DEV_MODE" default:"true"`
}

type database struct {
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     int    `env:"DB_PORT" default:"5432"`
	Username string `env:"DB_USER" default:"postgres"`
	Password string `env:"DB_PASS" default:""`
	Name     string `env:"DB_NAME" default:"cafe"`
	SSLMode  string `env:"DB_SSLMODE" default:"disable"`
}
