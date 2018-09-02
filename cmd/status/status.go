package main

import (
	"fmt"
	"os"
	"parking_lot/parking"
	"text/tabwriter"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		status := parking.Status()
		w := new(tabwriter.Writer)
		w.Init(os.Stderr, 0, 8, 0, '\t', 0)
		fmt.Fprintln(w, "Slot No.\tRegistration No\tColour")
		for _, parkingCar := range status {
			fmt.Fprintln(w, parkingCar)
		}
		fmt.Fprintln(w)
		w.Flush()
	}
}
