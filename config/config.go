package config

import (
	log "github.com/firekitz/fk-lib-log-go"
	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"strconv"
)

var LoadedConfig Config

type Config struct {
	ENV                 string `envconfig:"GO_ENV"`
	HTTP_PORT           string `envconfig:"HTTP_PORT"`
	GRPC_PORT           string `envconfig:"GRPC_PORT"`
	REDIS_HOST          string `envconfig:"REDIS_HOST"`
	REDIS_PORT          string `envconfig:"REDIS_PORT"`
	REDIS_PASSWORD      string `envconfig:"REDIS_PASSWORD"`
	REDIS_LOG_CONTAINER string `envconfig:"REDIS_LOG_CONTAINER"`
	LOG_SERVICE_NAME    string `envconfig:"LOG_SERVICE_NAME"`
	JWT_SECRET          string `envconfig:"JWT_SECRET"`
	PG_IAM_HOST         string `envconfig:"PG_IAM_HOST"`
	PG_IAM_USER         string `envconfig:"PG_IAM_USER"`
	PG_IAM_DATABASE     string `envconfig:"PG_IAM_DATABASE"`
	PG_IAM_PASSWORD     string `envconfig:"PG_IAM_PASSWORD"`
	PG_IAM_PORT         string `envconfig:"PG_IAM_PORT"`
}

func readEnv(cfg *Config) {
	err := envconfig.Process("", cfg)
	if err != nil {
		logrus.Fatalln("Fatal error config file", err)
	}
}

func localConfig(path string, config string) (err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(config)
	err = viper.ReadInConfig()
	if err != nil {
		logrus.Fatalln("Fatal error config file", err)
		return err
	}
	return nil
}

func LoadConfig(path string) (myConfig Config, err error) {
	if os.Getenv("GO_ENV") == "prod" {
		readEnv(&myConfig)
	} else if os.Getenv("GO_ENV") == "prod-local" {
		err := localConfig(path, "app.prod-local.env")
		if err != nil {
			return myConfig, err
		}
	} else if os.Getenv("GO_ENV") == "meta" {
		readEnv(&myConfig)
	} else if os.Getenv("GO_ENV") == "dev" {
		readEnv(&myConfig)
	} else if os.Getenv("GO_ENV") == "dev-local" {
		err := localConfig(path, "app.dev-local.env")
		if err != nil {
			return myConfig, err
		}
	} else if os.Getenv("GO_ENV") == "stage" {
		readEnv(&myConfig)
	} else if os.Getenv("GO_ENV") == "stage-local" {
		err := localConfig(path, "app.stage.env")
		if err != nil {
			return myConfig, err
		}
	} else {
		err := localConfig(path, "app.env")
		if err != nil {
			return myConfig, err
		}
	}
	//Handle errors reading the config file
	viper.AutomaticEnv()
	err = viper.Unmarshal(&myConfig)
	LoadedConfig = myConfig
	err = logSetup()
	if err != nil {
		return myConfig, err
	}

	return LoadedConfig, nil
}

func logSetup() error {
	port, _ := strconv.Atoi(LoadedConfig.REDIS_PORT)
	cfg := log.HookConfig{
		Host:     LoadedConfig.REDIS_HOST,
		Key:      LoadedConfig.REDIS_LOG_CONTAINER,
		Password: LoadedConfig.REDIS_PASSWORD,
		Port:     port,
		DB:       0,
		TTL:      0,
	}
	var label []string
	label = append(label, LoadedConfig.LOG_SERVICE_NAME, LoadedConfig.ENV)

	var err error
	err = log.Init(cfg, label)
	if err != nil {
		return err
	}
	return nil
}
