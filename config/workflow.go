package config

type Workflow struct {
	Name        string               `yaml:"name"`
	RunName     string               `yaml:"run-name"`
	On          map[string]OnTrigger `yaml:"on"`
	Permissions Permissions          `yaml:"permissions"`
	Env         map[string]string    `yaml:"env"`
	Defaults    Defaults             `yaml:"defaults"`
	Concurrency Concurrency          `yaml:"concurrency"`
	Jobs        map[string]Job       `yaml:"jobs"`
}

type OnTrigger struct {
	Types          []string                  `yaml:"types"`
	Branches       []string                  `yaml:"branches"`
	BranchesIgnore []string                  `yaml:"braches-ignore"`
	Paths          []string                  `yaml:"paths"`
	PathsIgnore    []string                  `yaml:"paths-ignore"`
	Tags           []string                  `yaml:"tags"`
	TagsIgnore     []string                  `yaml:"tags-ignore"`
	Cron           string                    `yaml:"cron"`
	Inputs         map[string]WorkflowInputs `yaml:"inputs"`
	Workflows      []string                  `yaml:"workflows"`
}

type WorkflowInputs struct {
	Description string   `yaml:"description"`
	Required    bool     `yaml:"required"`
	Default     string   `yaml:"default"`
	Type        string   `yaml:"type"`
	Options     []string `yaml:"options"`
}

type Permissions struct {
	Actions            string `yaml:"actions"`
	Checks             string `yaml:"checks"`
	Contents           string `yaml:"contents"`
	Deployments        string `yaml:"deployments"`
	Issues             string `yaml:"issues"`
	Discussions        string `yaml:"discussions"`
	Packages           string `yaml:"packages"`
	Pages              string `yaml:"pages"`
	PullRequests       string `yaml:"pull-requests"`
	RepositoryProjects string `yaml:"repository-projects"`
	SecurityEvents     string `yaml:"security-events"`
	Statuses           string `yaml:"statuses"`
}

type Defaults struct {
	Run Run `yaml:"run"`
}

type Run struct {
	Shell            string `yaml:"shell"`
	WorkingDirectory string `yaml:"working-directory"`
}

type Concurrency struct {
	Group            string `yaml:"group"`
	CancelInProgress bool   `yaml:"cancel-in-progress"`
}

type JobEnvironment struct {
	Name string `yaml:"name"`
	Url  string `yaml:"url"`
}

type Step struct {
	Id               string            `yaml:"id"`
	Name             string            `yaml:"name"`
	Uses             string            `yaml:"uses"`
	Run              string            `yaml:"run"`
	WorkingDirectory string            `yaml:"working-directory"`
	Shell            string            `yaml:"shell"`
	With             map[string]string `yaml:"with"`
	Environment      map[string]string `yaml:"environment"`
	ContinueOnError  bool              `yaml:"continue-on-error"`
	TimeOut          int               `yaml:"time-out"`
}

type Matrix struct {
	ValueList map[string]([]string)
	Include   []map[string]string `yaml:"include"`
	Exclude   []map[string]string `yaml:"exclude"`
}

type Strategy struct {
	Matrix      Matrix `yaml:"matrix"`
	FailFast    bool   `yaml:"fail-fast"`
	MaxParallel int    `yaml:"max-parallel"`
}

type Credentials struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

type Container struct {
	Image       string            `yaml:"image"`
	Credentials Credentials       `yaml:"credentials"`
	Env         map[string]string `yaml:"env"`
	Ports       []int             `yaml:"ports"`
	Volumes     []string          `yaml:"volumes"`
	Options     []string          `yaml:"options"`
}

type Services struct {
	Image       string            `yaml:"image"`
	Credentials Credentials       `yaml:"credentials"`
	Env         map[string]string `yaml:"env"`
	Ports       []int             `yaml:"ports"`
	Volumes     []string          `yaml:"volumes"`
	Options     []string          `yaml:"options"`
}

type Job struct {
	Id              string
	Name            string              `yaml:"name"`
	Permissions     Permissions         `yaml:"permissions"`
	Needs           []string            `yaml:"needs"`
	If              string              `yaml:"if"`
	RunsOn          string              `yaml:"runs-on"`
	JobEnvironment  JobEnvironment      `yaml:"environment"`
	Concurrency     Concurrency         `yaml:"concurrency"`
	Outputs         map[string]string   `yaml:"outputs"`
	Env             map[string]string   `yaml:"env"`
	Defaults        Defaults            `yaml:"defaults"`
	Steps           []Step              `yaml:"steps"`
	TimeOut         int                 `yaml:"time-out"`
	Strategy        Strategy            `yaml:"strategy"`
	ContinueOnError bool                `yaml:"continue-on-error"`
	Container       Container           `yaml:"container"`
	Services        map[string]Services `yaml:"services"`
	Uses            string              `yaml:"uses"`
	With            map[string]string   `yaml:"with"`
	Secrets         map[string]string   `yaml:"secrets"`
}
