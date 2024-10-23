package cmd

import (
	"fmt"
	"frate-go/config"
	"frate-go/ftemplate"
	"log"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

var InitCmd = &cobra.Command{
	Use:     "init(i)",
	Aliases: []string{"i", "init"},
	Short:   "Initialize a new Frate-go project",
	Run: func(cmd *cobra.Command, args []string) {
		projectName, _ := cmd.Flags().GetString("name")
		projectVersion, _ := cmd.Flags().GetString("projectVersion")
		cmakeVersion, _ := cmd.Flags().GetString("cmakeVersion")
		sourceDir, _ := cmd.Flags().GetString("sourceDir")
		buildDir, _ := cmd.Flags().GetString("buildDir")
		compiler, _ := cmd.Flags().GetString("compiler")
		includeDir, _ := cmd.Flags().GetString("includeDir")
		language, _ := cmd.Flags().GetString("language")
		languageVersion, _ := cmd.Flags().GetString("languageVersion")

		cfg := config.Config{
			CMakeVersion:    cmakeVersion,
			ProjectName:     projectName,
			ProjectVersion:  projectVersion,
			IncludeDir:      includeDir,
			BuildDir:        buildDir,
			Compiler:        compiler,
			Language:        language,
			LanguageVersion: languageVersion,
		}

		GenerateSource(sourceDir, language, &cfg)


		ftemplate.GenerateCmake(cfg)
    err := config.GenerateConfig(cfg)
    if  err != nil {
      log.Fatal("error generating config", err)
    }

	},
}

func init() {
	InitCmd.Flags().StringP("name", "n", "FrateProject", "Name of the project")
	InitCmd.Flags().StringP("default", "d", "FrateProject", "Name of the project")
	InitCmd.Flags().StringP("projectVersion", "p", "0.0.1", "version of c++")
	InitCmd.Flags().StringP("cmakeVersion", "V", "3.16", "CMake version to use")
	InitCmd.Flags().StringP("compiler", "c", "g++", "compiler to use")
	InitCmd.Flags().StringP("sourceDir", "s", "src", "source directory")
	InitCmd.Flags().StringP("buildDir", "b", "build", "build directory")
	InitCmd.Flags().StringP("language", "l", "cpp", "default language")
	InitCmd.Flags().StringP("languageVersion", "L", "20", "version of c++")
	InitCmd.Flags().StringP("includeDir", "I", "include", "Include Directories")
}

func GenerateSource(sourceDir string, lang string, cfg *config.Config) {
  err := os.Mkdir(sourceDir, 0700)
  if err != nil {
    log.Fatal("error creating directory", err)
  }
	var _ map[string]string
	fileName := os.Getenv("HOME") + "/main.cpp.gotmpl"
	tmpl, err := template.ParseFiles(fileName)
  if err != nil{
    fmt.Println(err)
  }
	var ext string
	if lang == "c" {
		ext = ".c"
	}
	if lang == "cpp" {
		ext = ".cpp"
	}
	fileName = sourceDir + "/main" + ext
	file, err := os.Create(fileName)
  cfg.SourceFiles = append(cfg.SourceFiles, fileName)
  if err != nil{
    fmt.Println(err)
  }
	err = tmpl.Execute(file, "")
  if err != nil {
    log.Fatal("error executing template", err)
  }
}
