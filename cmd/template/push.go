package template

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"frate-go/config"

	"github.com/spf13/cobra"

	"gopkg.in/yaml.v3"
)

var pushCmd = &cobra.Command{
	Use:   "push [template-directory]",
	Short: "Push a template to the server",
	Long:  `Push a template to the template server, including metadata.yaml and all template files.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		templateDir := args[0]

		metadata, err := config.LoadMetadata()
		if err != nil {
			fmt.Printf("Error loading metadata: %v\n", err)
			return
		}

		templateServerURL := metadata.Repos.Default

		err = pushTemplate(templateDir, templateServerURL)
		if err != nil {
			fmt.Printf("Error pushing template: %v\n", err)
		} else {
			fmt.Println("Template pushed successfully!")
		}
	},
}

func validateMetadata(metadataPath string) error {
	metadataFile, err := os.Open(metadataPath)
	if err != nil {
		return fmt.Errorf("unable to open metadata.yaml: %v", err)
	}
	defer metadataFile.Close()

	var metadata TemplateMetadata
	decoder := yaml.NewDecoder(metadataFile)
	err = decoder.Decode(metadata)

	if metadata.Name == "" || metadata.Version == "" || metadata.GitURL == "" {
		return fmt.Errorf("missing required fields in metadata")
	}

	return nil
}

func pushTemplate(templateDir, serverURL string) error {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	metadataPath := filepath.Join(templateDir, "metadata.yaml")
	metadataFile, err := os.Open(metadataPath)
	if err != nil {
		return fmt.Errorf("unable to open metadata.yaml: %v", err)
	}
  err = validateMetadata(metadataPath)
  if err != nil {
		return fmt.Errorf("unable to open metadata.yaml: %v", err)
  }
	defer metadataFile.Close()

	metadataWriter, err := writer.CreateFormFile("metadata", "metadata.yaml")
	if err != nil {
		return fmt.Errorf("unable to create form field for metadata: %v", err)
	}
	_, err = io.Copy(metadataWriter, metadataFile)
	if err != nil {
		return fmt.Errorf("unable to copy metadata.yaml to form: %v", err)
	}

	err = filepath.Walk(filepath.Join(templateDir, "templates"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return fmt.Errorf("unable to open file: %v", err)
			}
			defer file.Close()

			formFile, err := writer.CreateFormFile("files", filepath.Base(path))
			if err != nil {
				return fmt.Errorf("unable to create form field for file: %v", err)
			}

			_, err = io.Copy(formFile, file)
			if err != nil {
				return fmt.Errorf("unable to copy file to form: %v", err)
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("error walking template directory: %v", err)
	}

	writer.Close()

	req, err := http.NewRequest("POST", serverURL, body)
	if err != nil {
		return fmt.Errorf("unable to create request: %v", err)
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		respBody, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("server returned an error: %v - %s", resp.StatusCode, respBody)
	}

	return nil
}

func init() {
	TemplateCMD.AddCommand(pushCmd)
}
