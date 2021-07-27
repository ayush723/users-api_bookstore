package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

const (
	postgres_users_username = "ayush"
	postgres_users_password = "envy"
	postgres_users_host     = "localhost"
	postgres_users_schema   = "users_db"
)

var (
	Client *sql.DB
	username = os.Getenv(postgres_users_username)
	password = os.Getenv(postgres_users_password)
	host     = os.Getenv(postgres_users_host)
	schema   = os.Getenv(postgres_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("postgres://%s:%s@%s/%s",username,password,host,schema)


	var err error
	Client, err = sql.Open("postgres", dataSourceName)
	if err != nil {

		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("database successfully configured")

}
