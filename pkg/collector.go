package pkg

const (
	AsusCollectorType = "asus"
	MsiCollectorType  = "msi"
)

type Collector interface {
	SetCore()
	SetBrand()
	SetMemory()
	SetGraphicCard()
	GetComputer()
}

func GetCollector(collectorType string) *Collector {
	switch collectorType {
	case AsusCollectorType:
		return &AsusCollector{}
	case MsiCollectorType:
		return &MSICollector{}
	default:
		return nil
	}
}
