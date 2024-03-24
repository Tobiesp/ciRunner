package config

type DBConfig struct {
	Type     string   `yaml:"type"`
	FilePath string   `yaml:"file-path"`
	Username string   `yaml:"username"`
	Password string   `yaml:"password"`
	Server   string   `yaml:"server"`
	Protocol string   `yaml:"protocol"`
	Port     int      `yaml:"port"`
	DBName   string   `yaml:"db-name"`
	Options  []string `yaml:"options"`
}
