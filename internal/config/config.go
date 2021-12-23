package config

type Command struct {
	WorkDir string `yaml:"work_dir"`
	Command string `yaml:"command"`
}

type Project struct {
	Name     string     `yaml:"name"`
	ID       string     `yaml:"id"`
	Secret   string     `yaml:"secret"`
	Commands []*Command `yaml:"commands"`
}

type Config struct {
	AdminSecret string     `yaml:"admin_secret"`
	Port        int        `yaml:"port"`
	Projects    []*Project `yaml:"projects"`
}
