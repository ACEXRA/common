package main

import (
	"bufio"
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	dsn := "root:root@tcp(localhost:3306)/scim?parseTime=true"

	reader:=bufio.NewReader(os.Stdin)

	print("Enter up or down")
	flag,err:=reader.ReadString('\n');

	if err!=nil{
		log.Fatalf("Error reading input: %v",err)
	}

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
	switch flag{
	case "up":
		m.Up()
	case "down":
		m.Down()
	default:
		m.Up()
	}
    
	println("Migration ran successfully")
}