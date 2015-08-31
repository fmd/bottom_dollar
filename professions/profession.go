package professions

type Profession interface {
	Key() string
    String() string
	Salary() int
    Class() SocioEconomicClass
}