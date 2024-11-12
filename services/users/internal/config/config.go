package config

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
    Env string `yaml:"env" env-default:"local"`
    StoragePath string `yaml:"storage_path" env-required:"true"`
    DBName string `yaml:"db_name" env-required:"true"`
    AccessTokenTTL  time.Duration `yaml:"access_token_ttl" env-required:"true"`
    RefreshTokenTTL time.Duration `yaml:"refresh_token_ttl" env-required:"true"`
    JwtAccessSecret string `yaml:"jwt_access_secret" env-required:"true"`
    JwtRefreshSecret string `yaml:"jwt_refresh_secret" env-required:"true"`
    GRPC GRPCConfig `yaml:"grpc"`
    HTTP HTTPConfig `yaml:"http"`
    Email EmailConfig `yaml:"email"`
}

type GRPCConfig struct {
    Port    int     `yaml:"port" env-default:"44044"`
    Timeout time.Duration `yaml:"timeout"`
}

type HTTPConfig struct {
    Port int `yaml:"port" env-default:"8080"`
    Timeout time.Duration `yaml:"timeout"`
    IdleTimeout time.Duration `yaml:"idle_timeout"`
}

type EmailConfig struct {
    Port int `yaml:"port" env-default:"587"`
    User string `yaml:"user" env-required:"true"`
    Password string `yaml:"pswd" env-required:"true"`
    Host string `yaml:"host" env-required:"true"`
}

func MustLoad() *Config {
    path := fetchConfigPath()
    if path == "" {
        panic("config path is empty")
    }
    fmt.Println(path)

    if _, err := os.Stat(path); os.IsNotExist(err) {
        panic("config file does not exist")
    }

    var cfg Config

    if err := cleanenv.ReadConfig(path, &cfg); err != nil {
        panic("failed to read config: " + err.Error())
    }

    return &cfg
}

func fetchConfigPath() string {
    var res string

    flag.StringVar(&res, "config", "", "path to config file")
    flag.Parse()

    if res == "" {
        res = os.Getenv("CONFIG_PATH")
    }

    return res
}
