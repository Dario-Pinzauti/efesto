package db

import (
	"log"
	"script/query_script/storage"
)

type Form struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Database struct {
	DbName            string
	DbDefaultFileName string
}

var logger *log.Logger

func NewDatabase(dbName string, dbDefaultName string) *Database {
	db := &Database{
		DbName:            dbName,
		DbDefaultFileName: dbDefaultName,
	}

	return db
}

func Init() {
	logger = log.Default()
	logger.SetPrefix("database")
}

func CloseFile(d *Database, fileName string) {
	storage.CloseFile(fileName+"_"+d.DbName+".sql", d.DbDefaultFileName, logger)
}

func CreateNewFile(d *Database) {
	storage.CreateNewFile(d.DbDefaultFileName, logger)
}
