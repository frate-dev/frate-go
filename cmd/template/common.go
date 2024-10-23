package template


type TemplateMetadata struct {
	ID           int      `yaml:"id,omitempty" json:"id,omitempty"`
	Name         string   `yaml:"name" json:"name"`
	Version      string   `yaml:"version" json:"version"`
	InitScript   string   `yaml:"init_script,omitempty" json:"init_script,omitempty"`
	Description  string   `yaml:"description" json:"description"`
	Dependencies []string `yaml:"dependencies" json:"dependencies"`
	GitURL       string   `yaml:"git_url" json:"git_url"`
}
