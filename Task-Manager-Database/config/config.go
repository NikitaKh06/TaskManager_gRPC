package config

import (
	"context"
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/jackc/pgx/v5"
	"log"
	"os"
)

type ConfigDatabase struct {
	User     string `env:"DB_USER" env-default:"mydefaultuser"`
	Password string `env:"DB_PASSWORD" env-default:"mydefaultdpassword"`
	Name     string ` env:"DB_NAME" env-default:"mydefaultname"`
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     string `env:"DB_PORT" env-default:"5432"`
}

var Db *pgx.Conn

func ConfigureDatabase() error {
	var configDb ConfigDatabase

	err := cleanenv.ReadConfig("config/config.env", &configDb)
	if err != nil {
		return err
	}
	log.Println("Reading the config")

	db_url := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		configDb.User, configDb.Password, configDb.Host, configDb.Port, configDb.Name)

	connection, err := pgx.Connect(context.Background(), db_url)
	if err != nil {
		return err
	}

	log.Println("Connected to the database")
	Db = connection

	query := "" +
		"CREATE TABLE IF NOT EXISTS tasks (" +
		"id SERIAL PRIMARY KEY," +
		"name VARCHAR(50)," +
		"text VARCHAR(200)," +
		"done BOOLEAN DEFAULT FALSE" +
		");"

	_, err = Db.Exec(context.Background(), query)
	if err != nil {
		return err
	}

	log.Println("Created the table for tasks")

	return nil
}

var LogFile *os.File

func ConfigureLogger() error {
	logDirPath := "/app/logs"
	logFilePath := logDirPath + "/task-manager-database-logs.txt"

	err := os.MkdirAll(logDirPath, 0777)
	if err != nil {
		return err
	}

	logFile, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}

	LogFile = logFile

	log.SetOutput(LogFile)

	log.Println("Logger configured")

	return nil
}
