package professions

type Jobless struct {
	SalaryStore int
}

func NewJobless() *Jobless {
	d := &Jobless{}
	d.SalaryStore = 0
	return d
}

func (d *Jobless) String() string {
	return "vagrant"
}

func (d *Jobless) Key() string {
	return "jobless"
}

func (d *Jobless) Salary() int {
	return d.SalaryStore
}

func (d *Jobless) Class() SocioEconomicClass {
	return POOR
}
