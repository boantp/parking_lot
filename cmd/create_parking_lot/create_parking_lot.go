package main

import (
	"fmt"
	"os"
	"parking_lot/parking"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		cplValue := os.Args[1]
		cpl := parking.CreateParkingLot(cplValue)
		fmt.Println(cpl)
	}
}
