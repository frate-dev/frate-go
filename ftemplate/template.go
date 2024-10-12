package ftemplate

import (
	"log"
	"frate-go/config"
	"os"
	"text/template"
)

func GenerateCmake(cfg config.Config) {
	fileName := os.Getenv("HOME") + "/CMakeLists.txt.gotmpl"
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
    log.Fatal("error", err)
		return
	}
	file, err := os.Create("CMakeLists.txt")
	tmpl.Execute(file, cfg)
}
