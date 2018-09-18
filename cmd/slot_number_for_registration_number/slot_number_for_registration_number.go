package main

import (
	"fmt"
	"os"

	"github.com/boantp/parking_lot/parking"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		carPlateNumber := os.Args[1]
		rnWithColor := parking.SlotNumberForRegistrationNumber(carPlateNumber)
		fmt.Println(rnWithColor)
	}
}
