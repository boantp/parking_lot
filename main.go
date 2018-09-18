package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/user"
	"strconv"
	"strings"
	"text/tabwriter"

	"github.com/boantp/parking_lot/parking"
)

func main() {
	//Based on input file vs interactive
	arg := os.Args
	if len(arg) > 1 {
		fileName := os.Args[1]
		//input file
		if strings.Contains(fileName, "txt") {
			readFile := ReadFile(fileName)
			processReadFile(readFile)
		}
	} else if len(arg) == 1 {
		//interactive console
		processInteractive()
	}
}

//Process input read file
func processReadFile(readFile []string) {
	for _, value := range readFile {
		inputLine := strings.Split(value, "\n")
		ilString := inputLine[0]
		words := strings.Fields(ilString)
		for _, word := range words {
			switch word {
			case "create_parking_lot":
				cplValue := words[1]
				cpl := parking.CreateParkingLot(cplValue)
				fmt.Println(cpl)
			case "park":
				cpn := words[1]
				cc := words[2]
				park := parking.Park(cpn, cc)
				fmt.Println(park)
			case "leave":
				sn := words[1]
				sni, _ := strconv.Atoi(sn)
				leave := parking.Leave(sni)
				fmt.Println(leave)
			case "status":
				status := parking.Status()
				w := new(tabwriter.Writer)
				w.Init(os.Stderr, 0, 8, 0, '\t', 0)
				fmt.Fprintln(w, "Slot No.\tRegistration No\tColour")
				for _, parkingCar := range status {
					fmt.Fprintln(w, parkingCar)
				}
				fmt.Fprintln(w)
				w.Flush()
			case "registration_numbers_for_cars_with_colour":
				carColor := words[1]
				rnWithColor := parking.RegistrationNumbersForCarsWithColour(carColor)
				fmt.Println(rnWithColor)
			case "slot_numbers_for_cars_with_colour":
				carColor := words[1]
				rnWithColor := parking.SlotNumbersForCarsWithColour(carColor)
				fmt.Println(rnWithColor)
			case "slot_number_for_registration_number":
				carPlateNumber := words[1]
				rnWithColor := parking.SlotNumberForRegistrationNumber(carPlateNumber)
				fmt.Println(rnWithColor)
			}
		}
	}
}

func processInteractive() {
	// Get the current user.
	me, err := user.Current()
	if err != nil {
		panic(err)
	}

	// Get the current working directory.
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	// Set an environment variable.
	os.Setenv("SOME_VAR", "1")

	// Transfer stdin, stdout, and stderr to the new process
	// and also set target directory for the shell to start in.
	pa := os.ProcAttr{
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		Dir:   cwd,
	}

	// Start up a new shell.
	// Note that we supply "login" twice.
	// -fpl means "don't prompt for PW and pass through environment."
	fmt.Print(">> Starting a new interactive shell")
	proc, err := os.StartProcess("/usr/bin/login", []string{"login", "-fpl", me.Username}, &pa)
	if err != nil {
		panic(err)
	}

	// Wait until user exits the shell
	state, err := proc.Wait()
	if err != nil {
		panic(err)
	}

	// Keep on keepin' on.
	fmt.Printf("<< Exited shell: %s\n", state.String())
}

// ReadFile reads a file
func ReadFile(fileName string) []string {
	file, err := os.Open(getFile(fileName))
	defer file.Close()
	if err != nil {
		log.Fatalln(err)
	}

	br := bufio.NewReader(file)

	var lines []string

	for {
		// Includes the delimiter
		l, err := br.ReadString('\n')

		if err != nil && err != io.EOF {
			log.Fatalln(err)
		}

		// Trimming space to remove the delimiter at the end
		lines = append(lines, strings.TrimSpace(l))

		if err == io.EOF {
			break
		}
	}

	//log.Println(lines)
	return lines
}

//Get file from root
func getFile(fileName string) string {
	wd, _ := os.Getwd()

	if !strings.HasSuffix(wd, "file") {
		wd += ""
	}

	return wd + "/" + fileName
}
