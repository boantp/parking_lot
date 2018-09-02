package main

import (
	"fmt"
	"os"
	"parking_lot/parking"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 2 {
		sn := os.Args[1]
		sni, _ := strconv.Atoi(sn)
		leave := parking.Leave(sni)
		fmt.Println(leave)
	}
}
