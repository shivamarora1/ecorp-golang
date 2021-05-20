package main

import "fmt"

type Guitarist interface {
	// PlayGuitar prints out "Playing Guitar"
	// to the terminal
	PlayGuitar()
}

type BaseGuitarist struct {
	Name string
}

type AcousticGuitarist struct {
	Name string
}

func (b BaseGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Bass Guitar\n", b.Name)
}

func (b AcousticGuitarist) PlayGuitar() {
	fmt.Printf("%s plays the Acoustic Guitar\n", b.Name)
}

func main() {
	player1 := BaseGuitarist{"Paul"}
	player2 := AcousticGuitarist{"Ringo"}
	player3 := BaseGuitarist{"Bob"}
	player4 := AcousticGuitarist{"Nolan"}
	player5 := BaseGuitarist{"Chris"}
	player6 := AcousticGuitarist{"Jones"}

	gs := []Guitarist{}
	gs = append(gs, player1)
	gs = append(gs, player2)
	gs = append(gs, player3)
	gs = append(gs, player4)
	gs = append(gs, player5)
	gs = append(gs, player6)

	for _, obj := range gs {
		obj.PlayGuitar()
	}
}
