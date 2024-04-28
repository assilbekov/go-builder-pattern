package pkg

type MSICollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (m *MSICollector) SetCore() {
	m.Core = 6
}

func (m *MSICollector) SetBrand() {
	m.Brand = "MSI"
}

func (m *MSICollector) SetMemory() {
	m.Memory = 32
}

func (m *MSICollector) SetGraphicCard() {
	m.GraphicCard = 16
}

func (m *MSICollector) GetComputer() *Computer {
	return &Computer{
		Core:        m.Core,
		Brand:       m.Brand,
		Memory:      m.Memory,
		Monitor:     2,
		GraphicCard: m.GraphicCard,
	}
}
