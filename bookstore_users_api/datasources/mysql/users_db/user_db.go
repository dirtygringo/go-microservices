package users_db

import (
	"database/sql"
	"fmt"
	"log"

	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	mysql_users_username = "MYSQL_USERS_USERNAME"
	mysql_users_password = "MYSQL_USERS_PASSWORD"
	mysql_users_host     = "MYSQL_USERS_HOST"
	mysql_users_database = "MYSQL_USERS_DATABASE"
)

var (
	Client *sql.DB

	username = os.Getenv(mysql_users_username)
	password = os.Getenv(mysql_users_password)
	host     = os.Getenv(mysql_users_host)
	database = os.Getenv(mysql_users_database)
)

func init() {
	datasourceName :=
		fmt.Sprintf(
			"%s:%s@tcp(%s)/%s?charset=utf8",
			username, password, host, database)

	var err error
	Client, err := sql.Open("mysql", datasourceName)

	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
