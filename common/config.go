package common

type Config struct {
	Db struct {
		Dsn string
	}
	Http *http
}

type http struct {
	Addr string
}

func NewConfig() *Config {
	return &Config{
		Http: &http{
			Addr: ":18080",
		},
	}
}
