package config

type Config struct {
	Metrics []Metric `yaml:"metrics"`
}

type (
	MetricType string
)

type Metric struct {
	Name   string     `yaml:"name"`
	Type   MetricType `yaml:"type"`
	Help   string     `yaml:"help"`
	Labels []string   `yaml:"labels"`
	Query  string     `yaml:"query"`
}
