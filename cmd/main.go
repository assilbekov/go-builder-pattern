package main

import "builder/pkg"

func main() {
	asusCollector := pkg.GetCollector(pkg.AsusCollectorType)
	msiCollector := pkg.GetCollector(pkg.MsiCollectorType)
	factory := pkg.NewFactory(asusCollector)
	asusComputer := factory.CreateComputer()
	asusComputer.Print()
	factory.SetCollector(msiCollector)
	msiComputer := factory.CreateComputer()
	msiComputer.Print()
}
