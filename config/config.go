package config

var config *Config

type Config struct {
	Server Server `json:"server" mapstructure:"server"`
	Client Client `json:"client" mapstructure:"client"`
}

type Server struct {
	Host      string   `json:"host" mapstructure:"host"`
	Databases Database `json:"database" mapstructure:"database"`
	Router    []string `json:"route" mapstructure:"route"`
}

type Database struct {
	Postgres Postgres `json:"postgres" mapstructure:"postgres"`
	Redis    Redis    `json:"redis" mapstructure:"redis"`
	Mongo    Mongo    `json:"mongo" mapstructure:"mongo"`
	Mysql    Mysql    `json:"mysql" mapstructure:"mysql"`
}

type Postgres struct {
	Name   string `json:"name" mapstructure:"name"`
	Enable bool   `json:"enable" mapstructure:"enable"`
	URI    string `json:"uri" mapstructure:"uri"`
}

type Redis struct {
	Name   string `json:"name" mapstructure:"name"`
	Enable bool   `json:"enable" mapstructure:"enable"`
	URI    string `json:"uri" mapstructure:"uri"`
}

type Mongo struct {
	Name   string `json:"name" mapstructure:"name"`
	Enable bool   `json:"enable" mapstructure:"enable"`
	URI    string `json:"uri" mapstructure:"uri"`
}

type Mysql struct {
	Name   string `json:"name" mapstructure:"name"`
	Enable bool   `json:"enable" mapstructure:"enable"`
	URI    string `json:"uri" mapstructure:"uri"`
}

type Client struct {
	Host        string   `json:"host" mapstructure:"host"`
	AllowedCors []string `json:"allowed_cors" mapstructure:"allowed_cors"`
}

type Router struct {
	Route []string `json:"route" mapstructure:"route`
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
