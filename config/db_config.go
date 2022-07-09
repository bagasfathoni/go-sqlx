package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Config struct {
	DB *sqlx.DB
}

type dbEnv struct {
	Name string
	Host string
	Port string
	User string
	Pass string
	Driv string
}

func InitDbConfig() *Config {
	newConfig := new(Config)
	dsn := setDsn()

	conn, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connected to database...")
	}

	newConfig.DB = conn
	return newConfig
}

func CloseDB(db *sqlx.DB) {
	log.Println("Connection ended.")
	err := db.DB.Close()
	if err != nil {
		panic(err.Error())
	}
}

func setDsn() string {
	var dbEnv dbEnv
	dbEnv.Name = os.Getenv("DB_NAME")
	dbEnv.Host = os.Getenv("DB_HOST")
	dbEnv.Port = os.Getenv("DB_PORT")
	dbEnv.User = os.Getenv("DB_USER")
	dbEnv.Pass = os.Getenv("DB_PASS")
	dbEnv.Driv = os.Getenv("DB_DRIV")

	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable", dbEnv.Driv, dbEnv.User, dbEnv.Pass, dbEnv.Host, dbEnv.Port, dbEnv.Name)
	return dsn

}
