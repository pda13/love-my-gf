package config

import (
	"errors"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
	"reflect"
	"runtime"
)

const (
	EnvConfigPath string = "APP_CONFIG_PATH"
)

var (
	errConfigEnvEmpty    = errors.New("config path env is not set")
	errConfigFileInvalid = errors.New("config file has invalid format")
)

type Config struct {
	Env             string     `yaml:"env" env-required:"true"`
	ApplicationName string     `yaml:"application_name" env-required:"true"`
	GrpcServer      GrpcServer `yaml:"grpc_server" env-required:"true"`
}

type GrpcServer struct {
	Port string `yaml:"port" env:"GRPC_SERVER_PORT" env-required:"true"`
}

func Load() (*Config, error) {
	op := runtime.FuncForPC(reflect.ValueOf(Load).Pointer()).Name()

	pathToConfig := os.Getenv(EnvConfigPath)
	if pathToConfig == "" {
		return nil, fmt.Errorf("%s -> %w", op, errConfigEnvEmpty)
	}

	return loadByPath(pathToConfig)
}

func loadByPath(pathToConfig string) (*Config, error) {
	var config Config
	if err := cleanenv.ReadConfig(pathToConfig, &config); err != nil {
		return nil, fmt.Errorf("loadByPath failed -> %w; %w", errConfigFileInvalid, err)
	}

	return &config, nil
}
