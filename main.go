package main

import (
	"script/query_script/config"
	"script/query_script/gui"
)

func main() {

	err := config.InitConfig()

	preapareGui()
	if err != nil {
		gui.PrintError(err)
	}

	runGui()
}

func _main() {
	preapareGui()
	runGui()
}

func preapareGui() {
	gui.InitGui()
}

func runGui() {
	gui.FirstMenu()
}
