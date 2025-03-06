package store

type Config struct {
	// Path to database file
	DatabaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}