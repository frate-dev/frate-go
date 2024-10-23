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
  if err != nil {
    log.Fatal("couldn't create CMakeLists.txt", err)
  }
  err = tmpl.Execute(file, cfg)
  if err != nil {
    log.Fatal("error executing template", err)
  }

}
