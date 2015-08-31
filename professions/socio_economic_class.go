package professions

type SocioEconomicClass int

const (
	POOR SocioEconomicClass = iota
	WORKING
	MIDDLE
	RICH
)

func (s SocioEconomicClass) String() string {
	switch s {
	case POOR:
		return "poor"
	case WORKING:
		return "working class"
	case MIDDLE:
		return "middle class"
	case RICH:
		return "upper class"
	}
	return "nil"
}