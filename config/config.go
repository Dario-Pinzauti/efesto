package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	efestoerrors "script/query_script/efesto_errors"
)

type Config struct {
	Databases []string `json:"databases"`
}

type EfestoConfig struct {
	DefaultPath string `json:"default_path"`
}

var Logger *log.Logger
var Conf Config
var EfestoConf EfestoConfig

var ArrayScripts []string

func InitConfig() *efestoerrors.EfestoError {

	err := initEfestoConfig()

	if err != nil {
		return err
	}
	//

	dberr := initDbConfig()
	if dberr != nil {
		return dberr
	}

	//

	Logger = log.Default()
	Logger.SetPrefix("MAIN")

	templatesErr := initTemplateFolder()
	if templatesErr != nil {
		return templatesErr
	}
	return nil
}

func initEfestoConfig() *efestoerrors.EfestoError {
	efConf, err := os.Open("efesto_conf.json")
	if err != nil {
		return &efestoerrors.EfestoError{Text: "not found efesto_conf.json"}
	}

	defer efConf.Close()

	bv, err := ioutil.ReadAll(efConf)
	json.Unmarshal(bv, &EfestoConf)

	if err != nil {
		return &efestoerrors.EfestoError{Text: "format error in file efesto_conf.json"}
	}

	return nil
}

func initDbConfig() *efestoerrors.EfestoError {
	path := filepath.FromSlash(EfestoConf.DefaultPath + "/templates/config")
	c, err := os.Open(path)

	if err != nil {
		return &efestoerrors.EfestoError{Text: "format error in file " + path}

	}

	defer c.Close()

	byteValue, err := ioutil.ReadAll(c)
	json.Unmarshal(byteValue, &Conf)

	if err != nil {
		return &efestoerrors.EfestoError{Text: "format error in file " + path}

	}

	return nil
}

func initTemplateFolder() *efestoerrors.EfestoError {

	directory := EfestoConf.DefaultPath + "/templates/"

	// Open the directory.
	outputDirRead, err := os.Open(directory)

	if err != nil {
		return nil
	}

	// Call Readdir to get all files.
	outputDirFiles, _ := outputDirRead.Readdir(0)

	// Loop over files.
	for outputIndex := range outputDirFiles {

		outputFileHere := outputDirFiles[outputIndex]
		if !outputFileHere.IsDir() {
			continue
		}
		// Get name of file.
		outputNameHere := outputFileHere.Name()
		ArrayScripts = append(ArrayScripts, outputNameHere)

	}
	return nil

}
