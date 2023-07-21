package domain

type CofeeType struct {
	CofeeName     string `yaml:"name"`
	TimeToRefresh string `yaml:"time"`
	Limit         int    `yaml:"limit"`
}
