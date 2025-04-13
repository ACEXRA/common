package main

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	dsn := "root:root@tcp(localhost:3306)/scim?parseTime=true"

    db, err := sql.Open("mysql", dsn)
	if err!=nil{
		log.Fatalf("Database connection failed: %v",err)
	}
    driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil{
		log.Fatalf("Driver instance failed: %v",err)
	}
    m, err := migrate.NewWithDatabaseInstance(
        "file://migrate/migrations",
        "mysql", driver)
	if err != nil{
		log.Fatalf("New data instance failed: %v",err)
	}
    m.Up()
	println("Migration ran successfully")
}