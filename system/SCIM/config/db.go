package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func ConnectDB() *sql.DB{
	var err error
	dsn := "root:root@tcp(localhost:3306)/scim?parseTime=true"

	DB, err = sql.Open("mysql",dsn);

	if err!=nil{
		log.Fatal("Database connection failed: ",err)
	}

	if err= DB.Ping(); err!=nil{
		log.Fatal("Database unreachable: ",err)
	}

	println("Database connection success")
	return DB
}