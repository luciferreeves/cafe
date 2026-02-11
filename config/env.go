package config

import "time"

type server struct {
	Host           string `env:"SERVER_HOST" default:"localhost"`
	Port           int    `env:"SERVER_PORT" default:"8080"`
	AppSecret      string `env:"APP_SECRET" default:"mysecret"`
	AppName        string `env:"APP_NAME" default:"Shifoo's Cafe"`
	AppDescription string `env:"APP_DESCRIPTION" default:"A cozy place for close friends"`
	DevMode        bool   `env:"DEV_MODE" default:"true"`
}

type openid struct {
	DiscoveryURL string `env:"OPENID_DISCOVERY_URL" default:""`
	ClientID     string `env:"OPENID_CLIENT_ID" default:""`
	ClientSecret string `env:"OPENID_CLIENT_SECRET" default:""`
	CallbackURL  string `env:"OPENID_CALLBACK_URL" default:"http://localhost:8080/auth/callback"`
}

type database struct {
	Host     string `env:"DB_HOST" default:"localhost"`
	Port     int    `env:"DB_PORT" default:"5432"`
	Username string `env:"DB_USER" default:"postgres"`
	Password string `env:"DB_PASS" default:""`
	Name     string `env:"DB_NAME" default:"cafe"`
	SSLMode  string `env:"DB_SSLMODE" default:"disable"`
}

type session struct {
	CookieDomain   string        `env:"SESSION_COOKIE_DOMAIN" default:"localhost"`
	CookieName     string        `env:"SESSION_COOKIE_NAME" default:"cafe_session"`
	CookiePath     string        `env:"SESSION_COOKIE_PATH" default:"/"`
	CookieSameSite string        `env:"SESSION_COOKIE_SAME_SITE" default:"Lax"`
	CookieSecure   bool          `env:"SESSION_SECURE_COOKIE" default:"false"`
	CookieTimeout  time.Duration `env:"SESSION_TIMEOUT" default:"24h"`
}
