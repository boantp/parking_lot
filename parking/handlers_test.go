package parking

import (
	"testing"
)

func TestCreateParkingLot(t *testing.T) {
	total := CreateParkingLot("5")
	if total != "Created a parking lot with 5 slots" {
		t.Errorf("CreateParkingLot was incorrect, got: %s, want: %s.", total, "Created a parking lot with 5 slots")
	}
}

func TestPark(t *testing.T) {
	//reset
	CreateParkingLot("5")
	park := Park("KA-01-HH-1234", "White")
	if park != "Allocated slot number: 1" {
		t.Errorf("Park was incorrect, got: %s, want: %s.", park, "Allocated slot number: 1")
	}
}

func TestLeave(t *testing.T) {
	leave := Leave(1)
	if leave != "Slot number 1 is free" {
		t.Errorf("Leave was incorrect, got: %s, want: %s.", leave, "Slot number 1 is free")
	}
}

func TestStatus(t *testing.T) {
	CreateParkingLot("5")
	Park("KA-01-HH-1234", "White")
	all, _ := AllParkingCar("default", "")
	if len(all) != 1 {
		t.Errorf("Status was incorrect, got: %d, want: %d.", len(all), 1)
	}
}

func TestRegistrationNumbersForCarsWithColour(t *testing.T) {
	carPlateNumber := RegistrationNumbersForCarsWithColour("White")
	if carPlateNumber != "KA-01-HH-1234" {
		t.Errorf("RegistrationNumbersForCarsWithColour was incorrect, got: %s, want: %s.", carPlateNumber, "KA-01-HH-1234")
	}
}

func TestSlotNumbersForCarsWithColour(t *testing.T) {
	slotNumber := SlotNumbersForCarsWithColour("White")
	if slotNumber != "1" {
		t.Errorf("SlotNumbersForCarsWithColour was incorrect, got: %s, want: %s.", slotNumber, "1")
	}
}

func TestSlotNumberForRegistrationNumber(t *testing.T) {
	slotNumber := SlotNumberForRegistrationNumber("KA-01-HH-1234")
	if slotNumber != "1" {
		t.Errorf("Park was incorrect, got: %s, want: %s.", slotNumber, "1")
	}
}
