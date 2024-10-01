package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		fmt.Println("Not found locale .env file", err.Error())
	}
}

type Api struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	TokenSecret string `yaml:"token_secret"`
	Static      Static `yaml:"static"`
	Tls         Tls    `yaml:"tls"`
}

type Tls struct {
	Enable       bool   `yaml:"enable"`
	CertFilePath string `yaml:"cert_file_path"`
	KeyFilePath  string `yaml:"key_file_path"`
}

type Static struct {
	FilesPath string `yaml:"files_path"`
}

type Db struct {
	ConnStr string `yaml:"conn_str"`
}

type Config struct {
	Api Api `yaml:"api"`
	Db  Db  `yaml:"db"`
}

const (
	DefaultConfigPath = ""

	DefaultApiHost = "0.0.0.0"
	DefaultApiPort = "80"

	DefaultStaticPath = "web"

	DefaultDbConnStr = "file:track-balance.db?foreign_keys=on"
)

func New() (*Config, error) {
	cfg := &Config{
		Api{
			Host: DefaultApiHost,
			Port: DefaultApiPort,
			Static: Static{
				FilesPath: DefaultStaticPath,
			},
			Tls: Tls{
				Enable:       false,
				CertFilePath: "",
				KeyFilePath:  "",
			},
		},
		Db{
			ConnStr: DefaultDbConnStr,
		},
	}

	if key, ok := os.LookupEnv("RP_API_HOST"); ok {
		cfg.Api.Host = key
	}

	if key, ok := os.LookupEnv("RP_API_PORT"); ok {
		cfg.Api.Port = key
	}

	if key, ok := os.LookupEnv("RP_API_TOKEN_SECRET"); ok {
		cfg.Api.TokenSecret = key
	}

	if key, ok := os.LookupEnv("RP_STATIC_FILE_PATH"); ok {
		cfg.Api.Static.FilesPath = key
	}

	if key, ok := os.LookupEnv("RP_ENABLE_FILE_PATH"); ok {
		if keyBool, err := strconv.ParseBool(key); err == nil {
			cfg.Api.Tls.Enable = keyBool
		}
	}

	if key, ok := os.LookupEnv("RP_TLS_CERT_FILE_PATH"); ok {
		cfg.Api.Tls.CertFilePath = key
	}

	if key, ok := os.LookupEnv("RP_TLS_KEY_FILE_PATH"); ok {
		cfg.Api.Tls.KeyFilePath = key
	}

	if dbConn, ok := os.LookupEnv("RP_DATABASE_CONN_STRING"); ok {
		cfg.Db.ConnStr = dbConn
	}

	var err error

	switch {
	case *ConfigPathFlag != DefaultConfigPath:
		err = cleanenv.ReadConfig(*ConfigPathFlag, cfg)
	case len(DefaultConfigPath) > 0:
		err = cleanenv.ReadConfig(DefaultConfigPath, cfg)
	}

	if err != nil {
		return nil, err
	}

	return cfg, nil
}
