package pkg

type AsusCollector struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (a *AsusCollector) SetCore() {
	a.Core = 4
}

func (a *AsusCollector) SetBrand() {
	a.Brand = "Asus"
}

func (a *AsusCollector) SetMemory() {
	a.Memory = 16
}

func (a *AsusCollector) SetGraphicCard() {
	a.GraphicCard = 8
}

func (a *AsusCollector) GetComputer() *Computer {
	return &Computer{
		Core:        a.Core,
		Brand:       a.Brand,
		Memory:      a.Memory,
		Monitor:     1,
		GraphicCard: a.GraphicCard,
	}
}
