package cmd

import (
	"encoding/json"
	"fmt"
	"frate-go/config"
	"frate-go/ftemplate"
	"io"
	"log"
	"os"
	"path/filepath"
	"text/template"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

const frateDir = ".frate-go"
const templateSubDir = "templates"

var InitCmd = &cobra.Command{
	Use:     "init(i)",
	Aliases: []string{"i", "init"},
	Short:   "Initialize a new Frate-go project",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatal("Error: You must specify a directory to initialize the project in.")
		}
		projectDir := args[0]

		err := os.MkdirAll(projectDir, 0755)
		if err != nil {
			log.Fatal("Error creating project directory:", err)
			return
		}

		projectName, _ := cmd.Flags().GetString("name")
		projectVersion, _ := cmd.Flags().GetString("projectVersion")
		cmakeVersion, _ := cmd.Flags().GetString("cmakeVersion")
		sourceDir, _ := cmd.Flags().GetString("sourceDir")
		buildDir, _ := cmd.Flags().GetString("buildDir")
		compiler, _ := cmd.Flags().GetString("compiler")
		includeDir, _ := cmd.Flags().GetString("includeDir")
		language, _ := cmd.Flags().GetString("language")
		languageVersion, _ := cmd.Flags().GetString("languageVersion")
		repoflag, _ := cmd.Flags().GetString("repo")
		templateName, _ := cmd.Flags().GetString("template")

		metadata, err := config.LoadMetadata()
		if err != nil {
			log.Fatal("Error loading metadata: \n\t", err)
			return
		}

		url := metadata.Templates.Default.Url + "/list-templates"
		if repo := repoflag; repo != "" {
			for _, repo := range metadata.Templates.AdditionalRepos {
				if repo.Name == repoflag {
					url = repo.Url + "/list-templates"
				}
			}
		}

		data, err := Get(url)
		var templates []Template
		json.Unmarshal([]byte(data), &templates)

		cfg := config.Config{
			CMakeVersion:    cmakeVersion,
			ProjectName:     projectName,
			ProjectVersion:  projectVersion,
			IncludeDir:      includeDir,
			BuildDir:        buildDir,
			Compiler:        compiler,
			SourceDir:       sourceDir,
			Language:        language,
			LanguageVersion: languageVersion,
			Template:        templateName,
		}

		GenerateSource(projectDir, repoflag, &cfg)

		err = os.Chdir(projectDir)
		if err != nil {
			log.Fatal("Error changing to project directory:", err)
			return
		}

		ftemplate.GenerateCmake(&cfg)
		err = config.GenerateConfig(cfg)
		if err != nil {
			log.Fatal("Error generating config", err)
		}
	},
}

func init() {
	InitCmd.Flags().StringP("name", "n", "FrateProject", "Name of the project")
	InitCmd.Flags().StringP("default", "d", "FrateProject", "Name of the project")
	InitCmd.Flags().StringP("projectVersion", "p", "0.0.1", "Version of the project")
	InitCmd.Flags().StringP("cmakeVersion", "V", "3.16", "CMake version to use")
	InitCmd.Flags().StringP("compiler", "c", "g++", "Compiler to use")
	InitCmd.Flags().StringP("sourceDir", "s", "src", "Source directory")
	InitCmd.Flags().StringP("buildDir", "b", "build", "Build directory")
	InitCmd.Flags().StringP("language", "l", "cpp", "Default language")
	InitCmd.Flags().StringP("languageVersion", "L", "20", "Version of the language")
	InitCmd.Flags().StringP("includeDir", "I", "include", "Include directories")
	InitCmd.Flags().StringP("repo", "r", "", "Name of the repo to fetch template from")
	InitCmd.Flags().StringP("template", "t", "executable", "Name of the template to use")
}

func DirExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreateFrateDir() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	fratePath := filepath.Join(home, frateDir)
	err = os.MkdirAll(fratePath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create frate directory: %w", err)
	}

	return nil
}

func CreateTemplateDir(templateName string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	templatePath := filepath.Join(home, frateDir, templateSubDir)
	err = os.MkdirAll(templatePath, 0755)
	if err != nil {
		return "", fmt.Errorf("failed to create template directory: %w", err)
	}

	return templatePath, nil
}

func FetchAndStoreTemplate(templateName, gitUrl string) (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	templatePath := filepath.Join(home, frateDir, templateSubDir, templateName)
	if exists, _ := DirExists(templatePath); exists {
		fmt.Printf("Template %s already exists at %s\n", templateName, templatePath)
		return templatePath, nil
	}

	fmt.Println("Cloning template from repository...")
	RunCommand("git", "clone", gitUrl, templatePath)
	fmt.Println(gitUrl, templatePath) 

	if exists, _ := DirExists(templatePath); !exists {
		return "", fmt.Errorf("failed to clone template into %s", templatePath)
	}

	return templatePath, nil
}

func GenerateSource(baseSourceDir string, repoName string, cfg *config.Config) {
	err := os.MkdirAll(baseSourceDir, 0755)
	if err != nil {
		log.Fatal("Error creating base source directory:", err)
	}

	metadata, err := config.LoadMetadata()
	if err != nil {
		log.Fatal("Error loading metadata:", err)
		return
	}

	templatesUrl := metadata.Templates.Default.Url + "/list-templates"
	if repoName != "" {
		for _, repo := range metadata.Templates.AdditionalRepos {
			if repo.Name == repoName {
				templatesUrl = repo.Url + "/list-templates"
			}
		}
	}

	var templates []Template
	data, err := Get(templatesUrl)
	if err != nil {
		log.Fatal("Error fetching templates:", err)
	}
	json.Unmarshal([]byte(data), &templates)

	var url string
	for _, template := range templates {
		if template.Name == cfg.Template {
			url = template.GitUrl
			fmt.Println(cfg)
		}
	}

	templatePath, err := FetchAndStoreTemplate(cfg.Template, url)
	if err != nil {
		log.Fatal("Error fetching template:", err)
	}

	valuesYamlPath := filepath.Join(templatePath, "values.yaml")
	valuesFile, err := os.Open(valuesYamlPath)
	if err != nil {
		log.Fatal("Error opening values.yaml:", err)
	}
	defer valuesFile.Close()

	valuesMap := make(map[string]interface{})
	valuesYaml, err := io.ReadAll(valuesFile)
	if err != nil {
		log.Fatal("Error reading values.yaml:", err)
	}

	err = yaml.Unmarshal(valuesYaml, &valuesMap)
	if err != nil {
		log.Fatal("Error unmarshaling values.yaml:", err)
	}

	if frate, ok := valuesMap["frate"]; ok {
		cfg.FillStruct(frate.(map[string]interface{}))
	}
	err = filepath.Walk(filepath.Join(templatePath, "template"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if filepath.Base(path) == "CMakeLists.txt.gotmpl" {
			cmakeTemplateDir := filepath.Join(baseSourceDir, "CMakeTemplate")
			err := os.MkdirAll(cmakeTemplateDir, 0755)
			if err != nil {
				return fmt.Errorf("Error creating CMake template directory: %w", err)
			}
			CopyFile(path, filepath.Join(cmakeTemplateDir, "CMakeLists.txt.gotmpl"))
		}

		return nil
	})
	if err != nil {
		log.Fatal("Error walking through template files:", err)
	}

	fmt.Println("Generating source files...", cfg.SourceFiles)
	for _, file := range cfg.SourceFiles {
		fmt.Println("Generating file:", file)
		// Use the templateSourceDir path as a base
		os.MkdirAll(filepath.Join(baseSourceDir, cfg.SourceDir), 0700)
		fileName := filepath.Join(templatePath, "template", file)
		outputFile := filepath.Join(baseSourceDir, file[:len(file)-len(".gotmpl")])

		err := os.MkdirAll(filepath.Dir(outputFile), 0700)
		if err != nil {
			log.Fatal("Error creating output file directory:", err)
		}

		outFile, err := os.Create(outputFile)
		if err != nil {
			log.Fatal("Error creating output file:", err)
		}
		defer outFile.Close()

		tmpl, err := template.ParseFiles(fileName)
		if err != nil {
			log.Fatal("Error parsing template:", err)
		}
		err = tmpl.Execute(outFile, cfg)
		if err != nil {
			log.Fatal("Error executing template:", err)
		}
	}
}
