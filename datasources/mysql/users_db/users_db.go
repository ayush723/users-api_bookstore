package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ayush723/utils-go_bookstore/logger"
	"github.com/go-sql-driver/mysql"
)

// const (
// 	mysql_users_username = "ayush"
// 	mysql_users_password = "envy"
// 	mysql_users_host     = "127.0.0.1"
// 	// postgres_users_port     = 3306
// 	mysql_users_schema = "users_db"
// )

var (
	Client *sql.DB
	// username = os.Getenv(mysql_users_username)
	// password = os.Getenv(mysql_users_password)
	// host     = os.Getenv(mysql_users_host)
	// // port     = (os.Getenv(postgres_users_port))
	// schema = os.Getenv(mysql_users_schema)
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "ayush", "envy", "127.0.0.1", "users_db")
	// fmt.Println(dataSourceName)
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
