package domain

type CofeeType struct {
	CofeeName string `yaml:"name"`
	Time      string `yaml:"time"`
	Limit     int    `yaml:"limit"`
}
