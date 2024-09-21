package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"time"
)

type (
	Config struct {
		App   `yaml:"app"`
		HTTP  `yaml:"http"`
		Email `yaml:"email"`
		Api   `yaml:"api"`
		PG    `yaml:"pg"`
		JWT   `yaml:"jwt"`
		Hash  `yaml:"hash"`
	}

	App struct {
		Name string `yaml:"name" env:"APP_NAME"`
	}
	HTTP struct {
		Port            string        `yaml:"port" env:"HTTP_PORT"`
		ReadTimeout     time.Duration `yaml:"read_timeout" env:"HTTP_READ_TIMEOUT"`
		WriteTimeout    time.Duration `yaml:"write_timeout" env:"HTTP_WRITE_TIMEOUT"`
		ShutdownTimeout time.Duration `yaml:"shutdown_timeout" env:"HTTP_SHUTDOWN_TIMEOUT"`
	}

	Email struct {
		Topic string `yaml:"topic" env:"EMAIL_TOPIC"`
		Link  string `yaml:"link" env:"EMAIL_LINK"`
	}

	Api struct {
		Url string `yaml:"url" env:"API_URL"`
	}
	PG struct {
		ConnURI string `yaml:"conn_uri" env:"PG_URI"`
	}
	Hash struct {
		Salt string `yaml:"salt" env:"HASH_SALT"`
	}
	JWT struct {
		Singkey string `yaml:"singkey" env:"JWT_SIGNING_KEY"`
	}
)

func NewConfig() (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig("./config/config.yaml", cfg)
	if err != nil {
		return nil, errors.Wrap(err, "NewConfig: fail to read config.yaml file")
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, errors.Wrap(err, "NewConfig: fail to read env")
	}
	return cfg, err
}
