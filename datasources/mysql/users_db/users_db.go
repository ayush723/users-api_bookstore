package users_db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	usersDB *sql.DB
)

func init() {
	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "ayush", "envy", "127.0.0.1", "users_db")

	var err error
	usersDB, err = sql.Open("postgres", datasourceName)
	if err != nil {

		panic(err)
	}
	if err = usersDB.Ping();err !=nil {
		panic(err)
	}
	log.Panicln("database successfully configured")

}
