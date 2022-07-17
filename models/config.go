package models

type ServerConfig struct {
	Name string `env:"NAME_POSTGRES"`
	Host string `env:"HOST_POSTGRES"`
	Port string `env:"PORT_POSTGRES"`
	User string `env:"USER_POSTGRES"`
	Pass string `env:"PASS_POSTGRES"`

	AppName     string `env:"APP_NAME"`
	AppPort     string `env:"APP_PORT"`
	LogLevel    string `env:"LOG_LEVEL"`
	Environment string `env:"ENVIRONMENT"`
}
