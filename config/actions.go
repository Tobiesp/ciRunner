package config

type Action struct {
	Name        string                 `yaml:"name"`
	Author      string                 `yaml:"author"`
	Description string                 `yaml:"description"`
	Inputs      map[string]ActionInput `yaml:"inputs"`
	Outputs     map[string]struct {
		Description string `yaml:"description"`
		Value       string `yaml:"value"`
	} `yaml:"outputs"`
	Runs ActionRuns `yaml:"runs"`
}

type ActionRuns struct {
	Using          string            `yaml:"using"`
	Pre            string            `yaml:"pre"`
	PreIf          string            `yaml:"pre-if"`
	Main           string            `yaml:"main"`
	Post           string            `yaml:"post"`
	PostIf         string            `yaml:"post-if"`
	Steps          []Step            `yaml:"steps"`
	Image          string            `yaml:"image"`
	PreEntryPoint  string            `yaml:"pre-entrypoint"`
	EntryPoint     string            `yaml:"entrypoint"`
	PostEntryPoint string            `yaml:"post-entrypoint"`
	Env            map[string]string `yaml:"env"`
	Args           []string          `yaml:"args"`
	Branding       struct {
		Icon  string `yaml:"icon"`
		Color string `yaml:"color"`
	} `yaml:"branding"`
}

type ActionInput struct {
	Description       string `yaml:"description"`
	Required          bool   `yaml:"required"`
	Default           string `yaml:"default"`
	DeprecatedMessage string `yaml:"deprecatedMessage"`
}
