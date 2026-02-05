package configs

import (
	"log"
	"sync"
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	App struct {
		Name string `mapstructure:"NAME"`
		URL  string `mapstructure:"URL"`
	}

	DB struct {
		Postgres struct {
			Read struct {
				Host            string        `mapstructure:"HOST"`
				Port            string        `mapstructure:"PORT"`
				User            string        `mapstructure:"USER"`
				Password        string        `mapstructure:"PASSWORD"`
				Name            string        `mapstructure:"NAME"`
				SSLMode         string        `mapstructure:"SSLMODE"`
				MaxConnLifetime time.Duration `mapstructure:"MAX_CONNECTION_LIFETIME"`
				MaxIdleConn     int           `mapstructure:"MAX_IDLE_CONNECTION"`
				MaxOpenConn     int           `mapstructure:"MAX_OPEN_CONNECTION"`
			} `mapstructure:"READ"`
			Write struct {
				Host            string        `mapstructure:"HOST"`
				Port            string        `mapstructure:"PORT"`
				User            string        `mapstructure:"USER"`
				Password        string        `mapstructure:"PASSWORD"`
				Name            string        `mapstructure:"NAME"`
				SSLMode         string        `mapstructure:"SSLMODE"`
				MaxConnLifetime time.Duration `mapstructure:"MAX_CONNECTION_LIFETIME"`
				MaxIdleConn     int           `mapstructure:"MAX_IDLE_CONNECTION"`
				MaxOpenConn     int           `mapstructure:"MAX_OPEN_CONNECTION"`
			} `mapstructure:"WRITE"`
		} `mapstructure:"PG"`
	}

	Server struct {
		Env      string `mapstructure:"ENV"`
		LogLevel string `mapstructure:"LOG_LEVEL"`
		Port     string `mapstructure:"PORT"`
		Shutdown struct {
			CleanupPeriodSeconds int64 `mapstructure:"CLEANUP_PERIOD_SECONDS"`
			GracePeriodSeconds   int64 `mapstructure:"GRACE_PERIOD_SECONDS"`
		} `mapstructure:"SHUTDOWN"`
	}
}

var (
	conf Config
	once sync.Once
)

func Get() *Config {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Failed reading config file: %s", err)
	}

	once.Do(func() {
		log.Println("Service configuration initialized.")
		err = viper.Unmarshal(&conf)
		if err != nil {
			log.Fatalf("Failed to unmarshal config: %s", err)
		}
	})

	return &conf
}
