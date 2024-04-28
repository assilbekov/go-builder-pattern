package pkg

import "fmt"

type Computer struct {
	Core        int
	Brand       string
	Memory      int
	Monitor     int
	GraphicCard int
}

func (c *Computer) Print() {
	fmt.Println(
		"%s Core: [%d] Memory: [%d] GraphicCard: [%d] Monitor: [%d]\n",
		c.Brand, c.Core, c.Memory, c.GraphicCard, c.Monitor)
}
