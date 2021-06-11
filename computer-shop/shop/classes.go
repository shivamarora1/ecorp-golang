package shop

import (
	"errors"
	"fmt"
)

type Shop struct {
	AvailableModels []IComputer
}

type IComputer interface {
	Price() int
	ModelNum() int
	Name() string
	Info()
}

func (s *Shop) Open() {
	//Initialise different computer models and store them into availableModel list
	lenevo1 := WindowComputer{
		Computer: Computer{
			ModelNumber:     1,
			ModelName:       "Lenovo Ideapad S145 Ryzen",
			Cost:            29999,
			OperatingSystem: "Windows 10 HOME",
			Processor:       "INTEL 3 Dual Core",
			RAM:             4,
			HDD:             1000,
			ScreeSize:       15,
		}, Waranty: 2, TouchScreen: true,
	}
	lenevo2 := WindowComputer{
		Computer: Computer{
			ModelNumber:     2,
			ModelName:       "Lenovo IdeaPad 3 ",
			Cost:            36990,
			OperatingSystem: "Windows 10 HOME",
			Processor:       "CORE i3 10th GEN",
			RAM:             8,
			HDD:             256,
			ScreeSize:       14,
		}, Waranty: 5, TouchScreen: false,
	}

	acer := LinuxComputer{
		Computer: Computer{
			ModelNumber:     3,
			ModelName:       "Acer SF313-53",
			Cost:            60990,
			OperatingSystem: "Linux",
			Processor:       "CORE i5",
			RAM:             8,
			HDD:             512,
			ScreeSize:       14,
		}, BatteryType: "Lithium Ion", Distributions: "UBUNTU",
		Accesories: []string{"Bag", "Wireless Charging Pod"},
	}

	macAir := MacComputer{
		Computer: Computer{
			ModelNumber:     4,
			ModelName:       "Apple MGND3HNA MacBook Air",
			Cost:            92900,
			OperatingSystem: "macOS BigSur",
			Processor:       "Apple M1 Chip",
			RAM:             8,
			HDD:             256,
			ScreeSize:       13,
		}, Color: "Golden", TouchId: false, USBPorts: 3,
	}

	macPro := MacComputer{
		Computer: Computer{
			ModelNumber:     5,
			ModelName:       "Apple MacBook Pro",
			Cost:            179990,
			OperatingSystem: "macOS BigSur",
			Processor:       "Intel Core I7",
			RAM:             16,
			HDD:             512,
			ScreeSize:       16,
		}, Color: "Black", TouchId: false, USBPorts: 4,
	}

	s.AvailableModels = append(s.AvailableModels, lenevo1, lenevo2, acer, macAir, macPro)
}

func (s *Shop) List() {
	for _, obj := range s.AvailableModels {
		fmt.Printf("Model %d, Name: %s, Price %d\n", obj.ModelNum(), obj.Name(), obj.Price())
	}
}

func (s *Shop) Search(modelN int) (computer IComputer, err error) {
	for _, obj := range s.AvailableModels {
		if obj.ModelNum() == modelN {
			return obj, nil
		}
	}
	return nil, errors.New("Model not found")
}
