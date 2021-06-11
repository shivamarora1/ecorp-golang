package shop

import (
	"fmt"
)

type WindowComputer struct {
	Computer
	Waranty     int
	TouchScreen bool
}

func (wc WindowComputer) Info() {
	infoStr := fmt.Sprintf("Name: %s\n"+"Model num: %d\n"+"Cost: Rs. %d"+
		"OS: %s\n"+"Processor: %s\n"+"RAM: %d GB"+
		"HDD: %d GB\n"+"Screen Size: %d inches\n"+"Warranty: %d years\nTouchScreen: %v",
		wc.ModelName, wc.ModelNumber, wc.Price, wc.OperatingSystem, wc.Processor,
		wc.RAM, wc.HDD, wc.ScreeSize, wc.Waranty, wc.TouchScreen)
	fmt.Println(infoStr)
}

func (wc *WindowComputer) Reboot() {
	fmt.Printf("Rebooting operating system of computer")
}

func (wc *WindowComputer) InsertDisc() {
	fmt.Printf("Reading from digital disc")
}
