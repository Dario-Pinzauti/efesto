package main

import (
	"script/query_script/config"
	"script/query_script/gui"
)

func main() {

	preapareGui()
	err := config.InitConfig()
	if err != nil {

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
