package main

import (
	"fmt"

	"database/sql"
)

import _ "github.com/go-sql-driver/mysql"

var _connectUser string
var _connectPass string
var _address string
var _database string

func DoesDatabaseExist(address, username, password, database string) bool {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + address + ":3306)/" + database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return false
	}

	return true
}

func CreateDatabase(address, username, password, database string) error {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + address + ":3306)/"

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Creating database ", connectString, database)
	_, err = db.Exec("CREATE DATABASE " + database)
	if err != nil {
		panic(err)
	}

	return nil
}

func DoesTableExist(address, username, password, database, table string) bool {
	_connectUser = username
	_connectPass = password

	connectString := _connectUser + ":" + _connectPass + "@tcp(" + address + ":3306)/" + database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + database)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("SELECT 1 FROM " + table + " LIMIT 1")
	if err != nil {
		return false
	}

	return true
}

func DoCreateTable(address, username, password, database, table string) error {
	_connectUser = username
	_connectPass = password

	connectString := _connectUser + ":" + _connectPass + "@tcp(" + address + ":3306)/" + database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("USE " + database)
	if err != nil {
		return err
	}

	_, err = db.Exec("CREATE TABLE " + table + " (Airline varchar(2) NOT NULL, " +
		"  AirlineId varchar(32), " +
		"  SourceAirportCode varchar(32), " +
		"  SourceAirportId varchar(32), " +
		"  DestAirportCode varchar(32), " +
		"  DestAirportId varchar(32), " +
		"  Codeshare varchar(1), " +
		"  Stops varchar(1), " +
		"  Equipment varchar(32)) ")
	if err != nil {
		return err
	}

	return nil
}

func CreateTable(address, username, password, database, table string) error {
	_connectUser = username
	_connectPass = password
	_address = address
	_database = database

	if DoesDatabaseExist(address, username, password, database) == false {
		CreateDatabase(address, username, password, database)
	}

	if DoesTableExist(address, username, password, database, table) == false {
		DoCreateTable(address, username, password, database, table)
	}

	return nil
}

func GetAllRows(table string) FlightRoutes {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + _address + ":3306)/" + _database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + _database)
	if err != nil {
		panic(err)
	}

	var rv FlightRoutes

	rows, err := db.Query("SELECT * FROM " + table)
	for rows.Next() {
		var Airline string
		var AirlineId string
		var SourceAirportCode string
		var SourceAirportId string
		var DestAirportCode string
		var DestAirportId string
		var Codeshare string
		var Stops string
		var Equipment string

		err = rows.Scan(&Airline, &AirlineId, &SourceAirportCode, &SourceAirportId, &DestAirportCode, &DestAirportId, &Codeshare, &Stops, &Equipment)

		v := FlightRoute{Airline: Airline,
			AirlineId:         AirlineId,
			SourceAirportCode: SourceAirportCode,
			SourceAirportId:   SourceAirportId,
			DestAirportCode:   DestAirportCode,
			DestAirportId:     DestAirportId,
			Codeshare:         Codeshare,
			Stops:             Stops,
			Equipment:         Equipment,
		}

		rv = append(rv, v)
	}

	return rv
}

func GetRows(table, field, value string) FlightRoutes {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + _address + ":3306)/" + _database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + _database)
	if err != nil {
		panic(err)
	}

	var rv FlightRoutes

	query := "SELECT * FROM " + table + " WHERE `" + field + "`=\"" + value + "\""
	fmt.Println(query)
	rows, err := db.Query("SELECT * FROM " + table + " WHERE `" + field + "`=\"" + value + "\"")
	for rows.Next() {
		var Airline string
		var AirlineId string
		var SourceAirportCode string
		var SourceAirportId string
		var DestAirportCode string
		var DestAirportId string
		var Codeshare string
		var Stops string
		var Equipment string

		err = rows.Scan(&Airline, &AirlineId, &SourceAirportCode, &SourceAirportId, &DestAirportCode, &DestAirportId, &Codeshare, &Stops, &Equipment)

		v := FlightRoute{Airline: Airline,
			AirlineId:         AirlineId,
			SourceAirportCode: SourceAirportCode,
			SourceAirportId:   SourceAirportId,
			DestAirportCode:   DestAirportCode,
			DestAirportId:     DestAirportId,
			Codeshare:         Codeshare,
			Stops:             Stops,
			Equipment:         Equipment,
		}

		rv = append(rv, v)
	}

	return rv
}

func GetRowsMulti(table, field1, value1, field2, value2 string) FlightRoutes {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + _address + ":3306)/" + _database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	_, err = db.Exec("USE " + _database)
	if err != nil {
		panic(err)
	}

	var rv FlightRoutes

	rows, err := db.Query("SELECT * FROM " + table + " WHERE (`" + field1 + "`=\"" + value1 + "\" AND `" + field2 + "`=\"" + value2 + "\")")
	for rows.Next() {
		var Airline string
		var AirlineId string
		var SourceAirportCode string
		var SourceAirportId string
		var DestAirportCode string
		var DestAirportId string
		var Codeshare string
		var Stops string
		var Equipment string

		err = rows.Scan(&Airline, &AirlineId, &SourceAirportCode, &SourceAirportId, &DestAirportCode, &DestAirportId, &Codeshare, &Stops, &Equipment)

		v := FlightRoute{Airline: Airline,
			AirlineId:         AirlineId,
			SourceAirportCode: SourceAirportCode,
			SourceAirportId:   SourceAirportId,
			DestAirportCode:   DestAirportCode,
			DestAirportId:     DestAirportId,
			Codeshare:         Codeshare,
			Stops:             Stops,
			Equipment:         Equipment,
		}

		rv = append(rv, v)
	}

	return rv
}

func AddRow(table string, fields, values []string) error {
	connectString := _connectUser + ":" + _connectPass + "@tcp(" + _address + ":3306)/" + _database

	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return err
	}
	defer db.Close()

	_, err = db.Exec("USE " + _database)
	if err != nil {
		return err
	}

	insert := GetInsertCommand(table, fields, values)
	fmt.Println(insert)
	stmt, err := db.Prepare(insert)
	if err != nil {
		return err
	}

	_, err = stmt.Exec()
	if err != nil {
		pkField := 99
		update := GetUpdateCommand(table, pkField, fields, values)

		fmt.Println(update)
		stmt, err := db.Prepare(update)
		if err != nil {
			return err
		}

		_, err = stmt.Exec()
		if err != nil {
			panic(err)
		}
	}

	return nil
}

func GetInsertCommand(table string, fields, values []string) string {
	insert := "INSERT INTO " + table + "("
	for i, v := range fields {
		if i != 0 {
			insert += ", "
		}

		insert += v
	}

	insert += ") VALUES ("
	for i, v := range values {
		if i != 0 {
			insert += ", "
		}

		insert += "\"" + v + "\""
	}

	insert += ")"
	return insert
}

func GetUpdateCommand(table string, pkField int, fields, values []string) string {
	update := "UPDATE " + table + " set "
	for i, v := range fields {
		if i == pkField {
			continue
		}

		if i != 0 {
			update += ", "
		}

		update += v + "=\"" + values[i] + "\""
	}

	update += " WHERE "
	update += fields[pkField] + "=\"" + values[pkField] + "\""

	return update
}
