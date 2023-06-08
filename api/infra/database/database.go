package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

type DBConfig interface {
	ConnectDB() (*sql.DB, error)
	CloseDB(*sql.DB) error
}

type ConnectionDB struct {
	connect *sql.DB
	Driver  string
	URL     string
}

func (r *ConnectionDB) ConnectDB() (*sql.DB, error) {
	return sql.Open(r.Driver, r.URL)
}

func (r *ConnectionDB) CloseDB(db *sql.DB) error {
	return db.Close()
}

func NewDBConfig(driver string, url string) DBConfig {
	return &ConnectionDB{
		Driver: driver,
		URL:    url,
	}
}

func (r ConnectionDB) GetConnect() *sql.DB {
	return r.connect
}

func InstanceDB() *ConnectionDB {
	var (
		driver = "sqlite3"    //create a func global const and import os.Driver
		url    = "../test.db" //create a func global const and import os.URL
	)
	dbConfig := NewDBConfig(driver, url)
	connect, err := dbConfig.ConnectDB()
	if err != nil {
		panic(fmt.Sprintf("Don´t connection with DB: [%v]", err))
	}

	connect.SetMaxOpenConns(10)
	connect.SetMaxIdleConns(10)

	return &ConnectionDB{
		Driver:  driver,
		URL:     url,
		connect: connect,
	}

}

func CreateDatabase() {
	if _, err := os.Stat("./test.db"); os.IsNotExist(err) {
		file, err := os.Create("test.db")
		if err != nil {
			log.Fatal(err.Error())
		}
		file.Close()
		log.Println("sqlite-database.db created")
	}
}

func ScanStruct(rows *sql.Rows, dest interface{}) error {
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	if err := rows.Scan(scanArgs...); err != nil {
		return err
	}
	v := reflect.ValueOf(dest).Elem()
	for i, column := range columns {
		field := v.FieldByName(column)
		if !field.IsValid() {
			continue
		}
		if field.Type().Kind() == reflect.String {
			field.SetString(string(values[i]))
		} else {
			// Assume int
			val, err := strconv.ParseInt(string(values[i]), 10, 64)
			if err != nil {
				return err

			}
			field.SetInt(val)
		}
	}
	return nil
}
