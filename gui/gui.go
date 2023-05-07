package gui

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"script/query_script/config"
	efestoerrors "script/query_script/efesto_errors"
	fileactions "script/query_script/file_actions"
	"script/query_script/storage"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
)

var myApp fyne.App
var myWindow fyne.Window
var logger *log.Logger

func InitGui() {
	logger = log.Default()
	logger.SetPrefix("GUI")
	myApp = app.New()
	myWindow = myApp.NewWindow("Choice Widgets")
	FirstMenu()
	myWindow.Resize(fyne.NewSize(500, 500))
	myWindow.ShowAndRun()
}

func PrintError(e *efestoerrors.EfestoError) {
	lable := widget.NewLabel(e.Text)
	myWindow.SetContent(container.NewCenter(lable))
}

func FirstMenu() {
	var selectValue string

	button := widget.NewButton("next", func() {
		formWindow(selectValue)
	})
	buttonInitFile := widget.NewButton("init new file", func() {
		createNewFileWindow()
	})
	buttonCloseFile := widget.NewButton("close file", func() {
		closeFileWindow()
	})

	button.Disable()

	combo := widget.NewSelect(config.ArrayScripts, func(value string) {
		selectValue = value
		button.Enable()
	})

	myWindow.SetContent(container.NewVBox(combo, button, buttonInitFile, buttonCloseFile))
}

func createNewFileWindow() {
	fileactions.CreateNewFileAction()
	lable := widget.NewLabel("File Creato Correttamente.")

	myWindow.SetContent(container.NewVBox(lable, firstMenuButton()))
}

func closeFileWindow() {
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter file name...")

	btn := widget.NewButton("chiudi file", func() {
		fileactions.CloseFile(input.Text)
		lable := widget.NewLabel("File chiuso Correttamente.")
		myWindow.SetContent(container.NewVBox(lable, firstMenuButton()))
	})
	input.OnChanged = func(s string) {
		if s != "" {
			btn.Enable()
		} else {
			btn.Disable()
		}
	}
	myWindow.SetContent(container.NewVBox(input, btn))

}

func formWindow(value string) {

	formArray, dbArray := fileactions.PrepareScriptFolder(value)
	cont := container.NewVBox()
	//create form
	for _, v := range formArray {
		input := widget.NewEntry()
		input.SetPlaceHolder(v.Name)
		cont.Add(input)
	}

	btn := widget.NewButton("next", func() {
		obj := cont.Objects
		mymap := make(map[string]string)
		for _, v := range obj {
			input, ok := v.(*widget.Entry)
			if !ok {
				continue
			}
			mymap[input.PlaceHolder] = input.Text
		}

		for _, v := range dbArray {
			jsonFile, _ := os.Open(v.DbDefaultFileName)
			byteValue, _ := ioutil.ReadAll(jsonFile)
			for key, val := range mymap {

				byteValue = bytes.Replace(byteValue, []byte(stringForReplace(key)), []byte(val), -1)

			}
			storage.AppendInFile(string(byteValue), v.DbName, logger)

		}

		lable := widget.NewLabel(value + " inserito con successo")
		myWindow.SetContent(container.NewVBox(lable, firstMenuButton()))
	})

	cont.Add(btn)
	myWindow.SetContent(cont)
}

func stringForReplace(str string) string {
	return fmt.Sprintf("{{%s}}", str)
}

func firstMenuButton() fyne.Widget {
	return widget.NewButton("continue", func() {
		FirstMenu()
	})
}
