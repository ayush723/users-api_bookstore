package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/ayush723/utils-go_bookstore/logger"
	"github.com/go-sql-driver/mysql"
)

// const (
// 	mysql_users_username = "ayush"
// 	mysql_users_password = "wick"
// 	mysql_users_schema   = "users_db"
// 	mysql_users_host     = "127.0.0.1"
// )

var (
	Client *sql.DB
)

func init() {
	//setting environment variables
	// os.Setenv("mysql_users_username", mysql_users_username)
	// os.Setenv("mysql_users_password", mysql_users_password)
	// os.Setenv("mysql_users_schema", mysql_users_schema)
	// os.Setenv("mysql_users_host", mysql_users_host)

	//getting environment variables
	username := os.Getenv("mysql_users_username")
	password := os.Getenv("mysql_users_password")
	host := os.Getenv("mysql_users_host")
	schema := os.Getenv("mysql_users_schema")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, host, schema)
	fmt.Println(dataSourceName)
	// dataSourceName := "postgres://ayush:envy@localhost/users_db"
	var err error
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {

		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}
	mysql.SetLogger(logger.GetLogger())
	log.Println("database successfully configured")

}
