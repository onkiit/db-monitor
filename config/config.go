package config

var config *Config

type Config struct {
	Server Server `json:"server" mapstructure:"server"`
	Client Client `json:"client" mapstructure:"client"`
}

type Server struct {
	Host      string   `json:"host" mapstructure:"host"`
	Databases Database `json:"database" mapstructure:"database"`
}

type Database struct {
	Postgres Postgres `json:"postgres" mapstructure:"postgres"`
	Redis    Redis    `json:"redis" mapstructure:"redis"`
	Mongo    Mongo    `json:"mongo" mapstructure:"mongo"`
}

type Postgres struct {
	URI string `json:"uri" mapstructure:"uri"`
}

type Redis struct {
	URI string `json:"uri" mapstructure:"uri"`
}

type Mongo struct {
	URI string `json:"uri" mapstructure:"uri"`
}

type Client struct {
	Host        string   `json:"host" mapstructure:"host"`
	AllowedCors []string `json:"allowed_cors" mapstructure:"allowed_cors"`
}

func C(conf ...*Config) *Config {
	if len(conf) > 0 {
		config = conf[0]
	}
	return config
}

func init() {
	config = &Config{}
}