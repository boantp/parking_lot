package parking

import (
	"log"
	"strconv"
	"strings"
)

func CreateParkingLot(cplValue string) string {
	//Reset table
	TruncateTable()
	//insert into config with parking_lot_slot = 6
	config, err := PutConfig(cplValue)
	if err != nil {
		log.Fatalln(err)
	}
	pls := strconv.Itoa(config.ParkingLotSlot)

	//insert 6 rows data with parking_status = 0
	_, err = PutParkingLot(cplValue)
	if err != nil {
		log.Fatalln(err)
	}

	//Created a parking lot with 6 slots [OUTPUT]
	return "Created a parking lot with " + pls + " slots"
}

func Park(cpn string, cc string) string {
	//find the lowest id that parking_status = 0
	parkingLot, err := OneParkingLot()
	if err != nil {
		return "Sorry, parking lot is full"
	} else {
		//update table parking_lot
		_, err := UpdateParkingLot(parkingLot.SlotNumber, 1)
		if err != nil {
			log.Fatalln(err)
		}
		//insert table parking_car
		ppc, err := PutParkingCar(parkingLot.SlotNumber, cpn, cc)
		if err != nil {
			log.Fatalln(err)
		}
		sn := strconv.Itoa(ppc)
		return "Allocated slot number: " + sn
	}
}

func Leave(slotNumber int) string {
	//check from parking_car get slot_number
	parkingCar, err := OneParkingCar(slotNumber) //id_parking_car
	parkingCarId := parkingCar.IdParkingCar
	//update parking_lot with parking_status = 0 where slot_number
	_, err = UpdateParkingLot(slotNumber, 0)
	if err != nil {
		log.Fatalln(err)
	}
	//update parking_car with parking_status = 0 where slot_number and parking_status = 1
	_, err = UpdateParkingCar(parkingCarId, 0)
	if err != nil {
		log.Fatalln(err)
	}
	sn := strconv.Itoa(slotNumber)

	return "Slot number " + sn + " is free"
}

func Status() []string {
	all, err := AllParkingCar("default", "")
	if err != nil {
		log.Fatalln(err)
	}

	var status []string
	for _, value := range all {
		slotNumber := strconv.Itoa(value.SlotNumber)
		s := slotNumber + "\t" + value.CarPlateNumber + "\t" + value.CarColor
		status = append(status, s)
	}

	return status
}

func RegistrationNumbersForCarsWithColour(carColor string) string {
	all, err := AllParkingCar("color", carColor)
	if err != nil {
		log.Fatalln(err)
	}
	//KA-01-HH-1234, KA-01-HH-9999, KA-01-P-333
	var s string
	for _, value := range all {
		carPlateNumber := value.CarPlateNumber
		add := carPlateNumber + ", "
		s = s + add
	}
	status := trimSuffix(s, ", ")

	return status
}

func SlotNumbersForCarsWithColour(carColor string) string {
	all, err := AllParkingCar("color", carColor)
	if err != nil {
		log.Fatalln(err)
	}
	//1, 2, 4
	var s string
	for _, value := range all {
		slotNumber := strconv.Itoa(value.SlotNumber)
		add := slotNumber + ", "
		s = s + add
	}
	status := trimSuffix(s, ", ")

	return status
}

func SlotNumberForRegistrationNumber(carPlateNumber string) string {
	all, err := AllParkingCar("registration_number", carPlateNumber)
	if err != nil {
		log.Fatalln(err)
	}
	var slotNumber string
	if len(all) > 0 {
		for _, value := range all {
			slotNumber = strconv.Itoa(value.SlotNumber)
		}
	} else {
		slotNumber = "Not found"
	}
	return slotNumber
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}
