#PARKING_LOT

[![Build Status](https://travis-ci.org/boantp/parking_lot.svg?branch=master)](https://travis-ci.org/boantp/parking_lot)

## Getting started

Follow this step to run application in Golang + MySQL:
* Copy folder parking_lot into go/src, and copy into go/src/github.com for multiple package code
* Export sql from parking_lot.sql in the root folder in parking_lot
* Edit configuration database from parking_lot/config/db.php
```shell
DB, err = sql.Open("mysql", "USERNAME:PASSWORD@/DATABASE_NAME")
```
* Make It Runnable from the Command-Line for input file and interactive in folder 
parking_lot/cmd/create_parking_lot 
```shell
go build create_parking_lot.go
cp create_parking_lot /usr/local/bin
```
parking_lot/cmd/leave 
```shell
go build leave.go
cp leave /usr/local/bin
```
parking_lot/cmd/park 
```shell
go build park.go
cp park /usr/local/bin
```
parking_lot/cmd/registration_numbers_for_cars_with_colour 
```shell
go build registration_numbers_for_cars_with_colour.go
cp registration_numbers_for_cars_with_colour /usr/local/bin
```
parking_lot/cmd/slot_number_for_registration_number 
```shell
go build slot_number_for_registration_number.go
cp slot_number_for_registration_number /usr/local/bin
```
parking_lot/cmd/slot_numbers_for_cars_with_colour 
```shell
go build slot_numbers_for_cars_with_colour.go
cp slot_numbers_for_cars_with_colour /usr/local/bin
```
parking_lot/cmd/status 
```shell
go build status.go
cp status /usr/local/bin
```
parking_lot
```shell
go install
go build
```
* Unit Test in folder parking_lot/parking
```shell
go test -v
```
