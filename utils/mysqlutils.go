package utils

import (
	"database/sql"
	"log"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func CreateDatabaseIfNotExists(dsn string) error {

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("Failed to open MySQL")
		return err
	}
	defer db.Close()

	// create database
	stmt := `CREATE DATABASE IF NOT EXISTS beeblog`
	_, err = db.Exec(stmt)
	if err != nil {
		log.Println("Failed to create database")
		return err
	}

	return nil
}

func SQLinit() {
	dsn := "root:110119abc@tcp(127.0.0.1:3306)/?charset=utf8"
	err := CreateDatabaseIfNotExists(dsn)
	if err != nil {
		log.Fatalf("Create database failed: %v", err)
	}

	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:110119abc@tcp(127.0.0.1:3306)/beeblog?charset=utf8")
}
