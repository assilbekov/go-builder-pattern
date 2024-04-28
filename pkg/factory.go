package pkg

type Factory struct {
	Collector Collector
}

func NewFactory(collector Collector) *Factory {
	return &Factory{
		Collector: collector,
	}
}

func (f *Factory) SetCollector(collector Collector) {
	f.Collector = collector
}

func (f *Factory) CreateComputer() Computer {
	f.Collector.SetCore()
	f.Collector.SetBrand()
	f.Collector.SetMemory()
	f.Collector.SetGraphicCard()
	return *f.Collector.GetComputer()
}
