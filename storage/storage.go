package storage

import (
	"fmt"
	"log"
	"os"
)

func CreateNewFile(dbDefaultFileName string, logger *log.Logger) {
	f, err := os.Create(dbDefaultFileName)
	if err != nil {
		logger.Fatal(err)
	}
	defer f.Close()
}

func ArrayContains(sl []string, name string) bool {
	for _, value := range sl {
		if value == name {
			return true
		}
	}
	return false
}

func AppendInFile(s string, dbDefaultFileName string, logger *log.Logger) {
	f, err := os.OpenFile(dbDefaultFileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		logger.Fatal(err)
		return
	}
	_, err = fmt.Fprintln(f, s)
	if err != nil {
		logger.Fatal(err)
	}
	defer f.Close()

}

func ReadFile(dbDefaultFileName string, logger *log.Logger) []byte {
	f, err := os.ReadFile(dbDefaultFileName)

	if err != nil {
		logger.Fatal(err)
	}

	return f

}

func CloseFile(name string, dbDefaultFileName string, logger *log.Logger) {

	e := os.Rename(dbDefaultFileName, name)

	if e != nil {
		logger.Fatal(e)
	}

}
