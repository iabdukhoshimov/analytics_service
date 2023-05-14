package config

import (
	"fmt"
	"os"
	"time"

	env "github.com/Netflix/go-env"
	"github.com/joho/godotenv"
	logConfig "gitlab.com/voxe-analytics/pkg/logger/config"

	"gopkg.in/yaml.v3"
)

type AppMode string

const (
	DEVELOPMENT AppMode = "DEVELOPMENT"
	PRODUCTION  AppMode = "PRODUCTION"
)

type Config struct {
	Logging logConfig.Logging `yaml:"logging"`
	Project ProjectConfig     `yaml:"project"`
	Http    HttpConfig        `yaml:"http"`
	Storage StorageConfig     `yaml:"storage"`
	JWT     JWTConfig         `yaml:"jwt"`
	PSQL    PSQL
}

type ProjectConfig struct {
	Name                   string        `env:"PROJECT_NAME" yaml:"name"`
	Mode                   string        `env:"APPLICATION_MODE"`
	Version                string        `env:"APPLICATION_VERSION" yaml:"version"`
	Salt                   string        `env:"APP_SALT"`
	GracefulTimeoutSeconds int           `yaml:"gracefulTimeoutSeconds"`
	SwaggerEnabled         bool          `yaml:"swaggerEnabled"`
	FileUploadMaxMegabytes int           `yaml:"fileUploadMaxMegabytes"`
	Timeout                time.Duration `yaml:"timeout"`
}

type JWTConfig struct {
	JwtSecret             string        `env:"APPLICATION_JWT_SECRET"`
	RefreshSecret         string        `env:"APPLICATION_REFRESH_SECRET"`
	AccessTokenTTLMinutes time.Duration `yaml:"accessTokenTTLMinutes"`
	RefreshTokenTTLHours  time.Duration `yaml:"refreshTokenTTLHours"`
}

type HttpConfig struct {
	Host string `env:"HTTP_HOST" yaml:"host"`
	Port int    `env:"HTTP_PORT" yaml:"port"`
}

type StorageConfig struct {
	URI             string `env:"STORAGE_URI"`
	User            string `env:"STORAGE_MINIO_USER"`
	Password        string `env:"STORAGE_MINIO_PASSWORD"`
	AccessKeyID     string `env:"STORAGE_ACCESS_KEY_ID"`
	SecretAccessKey string `env:"STORAGE_SECRET_ACCESS_KEY"`
	BucketName      string `yaml:"bucketName"`
}

type PSQL struct {
	URI string `env:"PSQL_URI"`
}

func Load() *Config {
	cfg := Config{}

	err := godotenv.Load()
	if err != nil && !os.IsNotExist(err) {
		panic(err)
	}

	appMode := getAppMode()
	configPath, err := getConfigPath(appMode)
	if err != nil {
		panic(err)
	}

	file, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		panic(err)
	}

	_, err = env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		panic("unmarshal from environment error")
	}

	return &cfg
}

func getAppMode() AppMode {
	mode := AppMode(os.Getenv("APPLICATION_MODE"))
	if mode != PRODUCTION {
		mode = DEVELOPMENT
	}

	return mode
}

func getConfigPath(appMode AppMode) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	suffix := "dev"
	if appMode == PRODUCTION {
		suffix = "prod"
	}

	return fmt.Sprintf("%s/configs/app_config_%s.yaml", path, suffix), nil
}
