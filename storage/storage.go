package storage

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func FileExist(dbDefaultFileName string, logger *log.Logger) bool {
	filep := filepath.Dir(dbDefaultFileName)
	info, err := os.Stat(filep)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func CreateNewFile(dbDefaultFileName string, logger *log.Logger) {

	dirp := filepath.Dir(dbDefaultFileName)
	filep := filepath.Base(dbDefaultFileName)
	f, err := os.Create(filepath.Join(dirp, filep))
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

	dirp := filepath.Dir(dbDefaultFileName)
	filep := filepath.Base(dbDefaultFileName)
	f, err := os.OpenFile(filepath.Join(dirp, filep), os.O_APPEND|os.O_WRONLY, 0644)
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
	dirp := filepath.Dir(dbDefaultFileName)
	filep := filepath.Base(dbDefaultFileName)
	f, err := os.ReadFile(filepath.Join(dirp, filep))

	if err != nil {
		logger.Fatal(err)
	}

	return f

}

func CloseFile(name string, dbDefaultFileName string, logger *log.Logger) {
	dirp := filepath.Dir(dbDefaultFileName)
	filep := filepath.Base(dbDefaultFileName)
	e := os.Rename(filepath.Join(dirp, filep), name)

	if e != nil {
		logger.Fatal(e)
	}

}
