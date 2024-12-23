package main

import (
	"log/slog"
	"os"

	"github.com/joho/godotenv"
	"github.com/quanergyo/sigmaevolution/src/internal/repository/postgres"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		slog.Error("Error: init configs", err)
		os.Exit(1)
	}

	if err := godotenv.Load(); err != nil {
		slog.Error("Error: loading env variables", err)
		os.Exit(1)
	}

	db, err := postgres.NewDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		slog.Error("Error: failed to init db connection ", err)
		os.Exit(1)
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
