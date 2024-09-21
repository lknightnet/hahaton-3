package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/pkg/errors"
	"time"
)

type (
	Config struct {
		App  `yaml:"app"`
		HTTP `yaml:"http"`
		HTML `yaml:"html"`
		Mail `yaml:"mail"`
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

	HTML struct {
		Path string `yaml:"path" env:"HTML_PATH"`
	}
	Mail struct {
		From     string `yaml:"from" env:"MAIL_FROM"`
		Password string `yaml:"password" env:"MAIL_PASSWORD"`
		Host     string `yaml:"host" env:"MAIL_HOST"`
		Port     string `yaml:"port" env:"MAIL_PORT"`
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
