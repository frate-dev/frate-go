package template

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

func validateMetadata(metadataPath string)(TemplateMetadata, error) {
	metadataFile, err := os.Open(metadataPath)
	var metadata TemplateMetadata
	if err != nil {
		return metadata, fmt.Errorf("unable to open metadata.yaml: %v", err)
	}
	defer metadataFile.Close()

	data, err := io.ReadAll(metadataFile)
	if err != nil {
		return metadata,fmt.Errorf("error reading file")
	}
	err = yaml.Unmarshal(data, &metadata)
	fmt.Println("here metadata")
	if err != nil {
		return metadata, fmt.Errorf("error decoding metadata")
	}
	fmt.Println(metadata)

	if metadata.Name == "" || metadata.Version == "" || metadata.GitURL == "" {
		return metadata, fmt.Errorf("missing required fields in metadata")
	}

	return metadata, nil
}

func pushTemplate(templateDir, serverURL string) error {
	metadataPath := filepath.Join(templateDir, "metadata.yaml")
	metadataFile, err := os.Open(metadataPath)
	if err != nil {
		return fmt.Errorf("unable to open metadata.yaml: %v", err)
	}
	metadata, err := validateMetadata(metadataPath)
  if err != nil {
		return fmt.Errorf("unable to open metadata.yaml: %v", err)
  }
	defer metadataFile.Close()

	data, err := json.Marshal(metadata)
	if err != nil {
		return fmt.Errorf("unable to marshal to json: %v", err)
	}
	fmt.Println(string(data))
	req, err := http.NewRequest("POST", serverURL, bytes.NewBuffer(data))
	if err != nil {
		return fmt.Errorf("unable to create request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

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
