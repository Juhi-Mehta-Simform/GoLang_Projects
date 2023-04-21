package main

import (
	"fmt"
)

type printer interface {
	print()
}

type list []printer

func (l list) print() {
	if len(l) == 0 {
		fmt.Println("Sorry, Empty list.")
	}
	for _, it := range l {
		it.print()
	}
}

func (l list) discount(ratio float64) {
	for _, it := range l {
		g, isGame := it.(*game)
		if !isGame {
			continue
		}
		g.discount(ratio)
	}
}
