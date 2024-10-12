package dependency

type Dependency struct {
	Name         string       `yaml:"name"`
	GitURL       string       `yaml:"gitURL"`
	Tag          string       `yaml:"tag"`
	Link         string       `yaml:"link"`
	Dependencies []Dependency `yaml:"dependencies"`
}
