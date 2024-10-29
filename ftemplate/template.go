package ftemplate

import (
	"fmt"
	"frate-go/config"
	"log"
	"os"
	"path/filepath"
	"text/template"
)

func GenerateCmake(cfg *config.Config) {
	err := CollectSourceFiles(cfg)
	if err != nil {
		log.Fatalf("error collecting source files: %v", err)
	}
	projectDir, err := os.Getwd()
	if err != nil {
		log.Fatal("error fetching current directory:", err)
	}

	templateFilePath := filepath.Join(projectDir, "CMakeTemplate", "CMakeLists.txt.gotmpl")

	tmpl, err := template.ParseFiles(templateFilePath)
	if err != nil {
		log.Fatalf("error parsing template file %s: %v", templateFilePath, err)
	}

	outputFilePath := "CMakeLists.txt"
	file, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("couldn't create %s: %v", outputFilePath, err)
	}
	defer file.Close()

	err = tmpl.Execute(file, cfg)
	if err != nil {
		log.Fatalf("error executing template: %v", err)
	}
	fmt.Printf("Successfully generated %s\n", outputFilePath)
}

func CollectSourceFiles(cfg *config.Config) error {
	var sourceFiles []string

	err := filepath.Walk(cfg.SourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && (filepath.Ext(path) == ".cpp" || filepath.Ext(path) == ".c" || filepath.Ext(path) == ".h" || filepath.Ext(path) == ".hpp") {

			relPath, err := filepath.Rel(cfg.SourceDir, path)
			if err != nil {
				return fmt.Errorf("error getting relative path for %s: %v", path, err)
			}
			sourceFiles = append(sourceFiles, filepath.Join(cfg.SourceDir, relPath))
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking the source directory: %v", err)
	}

	cfg.SourceFiles = sourceFiles
	return nil
}
