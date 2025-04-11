package config

type AppConfig struct {
	PORT string `envconfig:"PORT" default:":8080"`
}
