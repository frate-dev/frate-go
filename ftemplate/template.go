package ftemplate

import (
	"fmt"
	"frate-go/config"
	"os"
	"text/template"
)

func GenerateCmake(cfg config.Config) {
	fileName := os.Getenv("HOME") + "/CMakeLists.txt.gotmpl"
  fmt.Println(cfg)
	tmpl, err := template.ParseFiles(fileName)
	if err != nil {
    fmt.Println("error", err)
		return
	}
	file, err := os.Create("CMakeLists.txt")
	tmpl.Execute(file, cfg)
  tmpl.Execute(os.Stdout, cfg)
}
