package main

import (
	conf "github.com/bagasfathoni/go-sqlx/config"
)

func main() {
	db := conf.InitDbConfig()
	defer conf.CloseDB(db.DB)
}
