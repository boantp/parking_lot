package main

import (
	"fmt"
	"os"

	"github.com/boantp/parking_lot/parking"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		cplValue := os.Args[1]
		cpl := parking.CreateParkingLot(cplValue)
		fmt.Println(cpl)
	}
}
