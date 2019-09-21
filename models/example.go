package models

import (
	"database/sql"
	"log"

	"github.com/go-gorp/gorp"
	_ "github.com/mattn/go-sqlite3"

	"../utils"
)

// Example struct
type Example struct {
	ID        int
	Name, URL string
}

// GetDatabase returns a gorp.DbMap that can be used to query data in an sql database
//   Allows for an entity framework (asp.net)-like interface to quering server data
func GetDatabase() *gorp.DbMap {
	db, err := sql.Open(utils.GetConfig().SQLDriver, utils.GetConfig().SQL)
	if err != nil {
		log.Fatal(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}

	// Add a table to the map, maps the Example row to the table "example",
	//   SetKeys tells us what the primary key is
	dbmap.AddTableWithName(Example{}, "example").SetKeys(true, "ID")

	err = dbmap.CreateTablesIfNotExists()
	if err != nil {
		log.Fatal(err)
	}

	// Insert example item
	// dbmap.Insert(&Example{Name: "test"})

	return dbmap
}
