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

var config utils.Config = utils.GetConfig()

// GetDatabase returns a gorp.DbMap that can be used to query data in an sql database
//   Allows for an entity framework (asp.net)-like interface to quering server data
func GetDatabase() *gorp.DbMap {
	db, err := sql.Open(config.SQLDriver, config.SQL)
	if err != nil {
		log.Fatal(err)
	}

	// change gorp.SqliteDialect to what server driver your using
	//    https://github.com/go-gorp/gorp#database-drivers
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
