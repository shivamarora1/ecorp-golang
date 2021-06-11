package shop

import (
	"fmt"
)

type MacComputer struct {
	Computer
	Color    string
	TouchId  bool
	USBPorts int
}

func (mc MacComputer) Info() {
	infoStr := fmt.Sprintf("Name: %s\n"+"Model num: %d\n"+"Cost: Rs. %d"+
		"OS: %s\n"+"Processor: %s\n"+"RAM: %d GB"+
		"HDD: %d GB\n"+"Screen Size: %d inches\n"+"Color: %s \nUSB Ports: %v\nTouchId: %v",
		mc.ModelName, mc.ModelNumber, mc.Cost, mc.OperatingSystem, mc.Processor,
		mc.RAM, mc.HDD, mc.ScreeSize, mc.Color, mc.USBPorts, mc.TouchId)
	fmt.Println(infoStr)
}

func (mc MacComputer) ApplePlus() {
	fmt.Println("Apple + service is available for Rs. 8500")
}
