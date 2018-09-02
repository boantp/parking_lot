package parking

import (
	"strconv"

	"github.com/parking_lot/config"
)

type Config struct {
	Id             int
	ParkingLotSlot int
}

type ParkingLot struct {
	IdParkingLot  int64
	SlotNumber    int
	ParkingStatus int
}

type ParkingCar struct {
	IdParkingCar   int
	SlotNumber     int
	CarPlateNumber string
	CarColor       string
	ParkingStatus  int
	DateTime       string
}

func PutConfig(cplValue string) (Config, error) {
	configData := Config{}
	i, _ := strconv.Atoi(cplValue)
	configData.ParkingLotSlot = i

	// insert values
	sqlStr := "INSERT INTO config(parking_lot_slot) VALUES (?)"
	//prepare the statement
	stmt, err := config.DB.Prepare(sqlStr)
	checkErr(err)

	//format all vals at once
	_, err = stmt.Exec(configData.ParkingLotSlot)
	checkErr(err)
	if err != nil {
		return configData, err
	}

	return configData, nil
}

func PutParkingLot(cplValue string) (ParkingLot, error) {
	pl := ParkingLot{}
	i, _ := strconv.Atoi(cplValue)

	for j := 1; j <= i; j++ {
		// insert values
		sqlStr := "INSERT INTO parking_lot(slot_number) VALUES (?)"
		//prepare the statement
		stmt, err := config.DB.Prepare(sqlStr)
		checkErr(err)

		//format all vals at once
		res, err := stmt.Exec(j)
		checkErr(err)
		id, err := res.LastInsertId()
		pl.IdParkingLot = id
	}

	return pl, nil
}

func PutParkingCar(slotNumber int, cpn string, cc string) (int, error) {
	// insert values
	sqlStr := "INSERT INTO parking_car(slot_number,car_plate_number,car_color,parking_status) VALUES (?,?,?,?)"
	//prepare the statement
	stmt, err := config.DB.Prepare(sqlStr)
	checkErr(err)

	//format all vals at once
	_, err = stmt.Exec(slotNumber, cpn, cc, 1)
	checkErr(err)
	if err != nil {
		return slotNumber, err
	}

	return slotNumber, err
}

func UpdateParkingLot(slotNumber int, parkingStatus int) (int, error) {
	// update
	stmt, err := config.DB.Prepare("update parking_lot set parking_status=? where slot_number=?")
	checkErr(err)

	_, err = stmt.Exec(parkingStatus, slotNumber)
	checkErr(err)
	if err != nil {
		return slotNumber, err
	}

	return slotNumber, err
}

func UpdateParkingCar(parkingCarId int, parkingStatus int) (int, error) {
	// update
	stmt, err := config.DB.Prepare("update parking_car set parking_status=? where id_parking_car=?")
	checkErr(err)

	_, err = stmt.Exec(parkingStatus, parkingCarId)
	checkErr(err)
	if err != nil {
		return parkingCarId, err
	}

	return parkingCarId, err
}

func OneParkingLot() (ParkingLot, error) {
	pl := ParkingLot{}
	// query
	row := config.DB.QueryRow("SELECT * FROM parking_lot WHERE parking_status = 0 ORDER BY id_parking_lot ASC LIMIT 1")
	err := row.Scan(&pl.IdParkingLot, &pl.SlotNumber, &pl.ParkingStatus)
	if err != nil {
		return pl, err
	}

	return pl, nil
}

func OneParkingCar(slotNumber int) (ParkingCar, error) {
	pc := ParkingCar{}
	// query
	rows, err := config.DB.Query("SELECT * FROM parking_car WHERE parking_status = 1 AND slot_number = ?", slotNumber)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&pc.IdParkingCar, &pc.SlotNumber, &pc.CarPlateNumber, &pc.CarColor, &pc.ParkingStatus, &pc.DateTime)
		checkErr(err)
	}

	if err != nil {
		return pc, err
	}

	return pc, nil
}

func AllParkingCar(key string, value string) ([]ParkingCar, error) {
	var where string
	if key == "color" {
		where = " AND car_color='" + value + "'"
	} else if key == "registration_number" {
		where = " AND car_plate_number='" + value + "'"
	} else if key == "default" {
		where = ""
	}
	//query
	query := "SELECT * FROM parking_car where parking_status=1" + where
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cars := make([]ParkingCar, 0)
	for rows.Next() {
		car := ParkingCar{}
		err := rows.Scan(&car.IdParkingCar, &car.SlotNumber, &car.CarPlateNumber, &car.CarColor, &car.ParkingStatus, &car.DateTime) // order matters
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return cars, nil
}

func TruncateTable() {
	// query
	_, err := config.DB.Query("TRUNCATE TABLE config")
	checkErr(err)
	_, err = config.DB.Query("TRUNCATE TABLE parking_lot")
	checkErr(err)
	_, err = config.DB.Query("TRUNCATE TABLE parking_car")
	checkErr(err)

	return
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
