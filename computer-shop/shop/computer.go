package shop

// base class
type Computer struct {
	ModelNumber     int
	ModelName       string
	Cost            int
	OperatingSystem string
	Processor       string
	RAM             int
	HDD             int
	ScreeSize       int //In inches
}

func (c Computer) ModelNum() int {
	return c.ModelNumber
}

func (c Computer) Name() string {
	return c.ModelName
}

func (c Computer) Price() int {
	return c.Cost
}
