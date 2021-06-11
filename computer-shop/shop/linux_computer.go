package shop

import (
	"fmt"
	"strings"
)

type LinuxComputer struct {
	Computer
	BatteryType   string
	Distributions string
	Accesories    []string
}

func (lc LinuxComputer) Info() {
	infoStr := fmt.Sprintf("Name: %s\n"+"Model num: %d\n"+"Cost: Rs. %d"+
		"OS: %s\n"+"Processor: %s\n"+"RAM: %d GB"+
		"HDD: %d GB\n"+"Screen Size: %d inches\n"+"Battery Type: %s \nAccessories: %v",
		lc.ModelName, lc.ModelNumber, lc.Price, lc.OperatingSystem, lc.Processor,
		lc.RAM, lc.HDD, lc.ScreeSize, lc.BatteryType, strings.Join(lc.Accesories, ","))
	fmt.Println(infoStr)
}
