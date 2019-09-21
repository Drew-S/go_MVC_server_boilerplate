package models

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"

	"../utils"
)

type Example struct {
	Name, Url string
}

func GetDatabase() *gorp.DbMap {
	db, err := sql.Open(utils.GetConfig().SQLDriver, utils.GetConfig().SQL)
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	dbmap.AddTableWithName(Example{}, "example").SetKeys(true, "Id")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	// Insert example item
	// dbmap.Insert(&Example{Name: "test"})

	return dbmap
}
