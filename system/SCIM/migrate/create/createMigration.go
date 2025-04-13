package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	println("Enter the table name: ")

	tableName, err:=reader.ReadString('\n');
	if err!=nil{
		log.Fatalf("Error in reading input: %v",err)
	}

	tableName = strings.TrimSpace(tableName)

	if tableName ==""{
		log.Fatal("Table name cannot be empty")
	}

	timestamp:= time.Now().Format("20060102150405")
	upFile := fmt.Sprintf("%s_%s.up.sql", timestamp, tableName)
	downFile := fmt.Sprintf("%s_%s.down.sql", timestamp, tableName)

	migrationDir := "./migrate/migrations"

	//create directory if it doesnt exist
	if _,err:=os.Stat(migrationDir);os.IsNotExist(err){
		err := os.MkdirAll(migrationDir,os.ModeAppend)
		if err!=nil{
			log.Fatalf("Error creating the directory: %v",err)
		}
	}

	//check the table name exist
	existingFile,_ := filepath.Glob(filepath.Join(migrationDir, "*_"+tableName+".up.sql"))
	if(len(existingFile)>0){
		log.Fatalf("A migration for '%s' already exists: %s",tableName,existingFile[0])
	}

	//create migration file in the directory
	upFilePath := filepath.Join(migrationDir,upFile)
	downFilePath:= filepath.Join(migrationDir,downFile)

	_,err=os.Create(upFilePath)
	if err!=nil{
		log.Fatalf("Error creating up file: %v",err)
	}
	_,err=os.Create(downFilePath)
	if err!=nil{
		log.Fatalf("Error creating down file: %v",err)
	}

}