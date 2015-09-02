package people

type Profession interface {
	Key() string
	String() string
	Salary() int
	Class() SocioEconomicClass
}
