package common

type Config struct {
	Db struct {
		Driver string
		Dsn    string
	}
	Http  http
	Debug bool
}

type http struct {
	Addr string
}

func NewConfig() *Config {
	return &Config{
		Http: http{
			Addr: ":18080",
		},
	}
}
