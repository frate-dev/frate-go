package template


type TemplateMetadata struct {
    ID          int      `json:"id,omitempty"`
    Name        string   `json:"name"`
    Version     string   `json:"version"`
    Description string   `json:"description"`
    Dependencies []string `json:"dependencies"`
    GitURL      string   `json:"git_url"`
    CreatedAt   string   `json:"created_at,omitempty"`
    License     string   `json:"license,omitempty"`
}

