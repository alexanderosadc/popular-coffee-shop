package domain

type MembershipType struct {
	MembershipName string
	TypesOfCofee   []CofeeType `yaml:"cofee_types"`
}
