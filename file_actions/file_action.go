package fileactions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"script/query_script/config"
	"script/query_script/db"
	"script/query_script/storage"
)

func CreateNewFileAction() {

	for _, v := range config.Conf.Databases {
		storage.CreateNewFile(v, config.Logger)
	}

}

func FileActionWasCreated() bool {
	var isCreated bool = false
	for _, v := range config.Conf.Databases {
		isCreated = storage.FileExist(v, config.Logger)
		if isCreated {
			return true
		}
	}
	return false

}

func CloseFile(name string) {
	for _, v := range config.Conf.Databases {
		storage.CloseFile(name+"_"+v+".sql", v, config.Logger)
	}
}

func PrepareScriptFolder(namePath string) ([]db.Form, []*db.Database) {
	outputDirRead, _ := os.Open(config.EfestoConf.DefaultPath + "/templates/" + namePath)
	var dbArr []*db.Database
	var formArray []db.Form
	// Call Readdir to get all files.
	defer outputDirRead.Close()
	outputDirFiles, _ := outputDirRead.Readdir(0)

	// Loop over files.
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		// Get name of file.
		outputNameHere := outputFileHere.Name()
		if outputFileHere.Name() == "form" {
			formFile, _ := os.Open(config.EfestoConf.DefaultPath + "/templates/" + namePath + "/" + outputNameHere)
			defer formFile.Close()
			byteValue, _ := ioutil.ReadAll(formFile)
			json.Unmarshal(byteValue, &formArray)
			continue
			//creare oggetto lista db e form
		}

		dbArr = append(dbArr, db.NewDatabase(outputFileHere.Name(), config.EfestoConf.DefaultPath+"/templates/"+namePath+"/"+outputFileHere.Name()))

		// Print name.
		fmt.Println(outputNameHere)
	}
	return formArray, dbArr
}
