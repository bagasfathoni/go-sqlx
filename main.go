package main

import (
	conf "github.com/bagasfathoni/go-sqlx/config"
	"github.com/jmoiron/sqlx"
)

func main() {
	db := conf.InitDbConfig()

	defer func(db *sqlx.DB) {
		err := db.DB.Close()
		if err != nil {
			panic(err.Error())
		}
	}(db.DB)
}
