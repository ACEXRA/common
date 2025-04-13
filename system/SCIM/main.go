package main

import (
	"github.com/acexra/scim/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config.ConnectDB();
}
