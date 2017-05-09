package dbhandler

import (
	"database/sql"
	"github.com/chandanchowdhury/HSPC/models"
	"log"
)

/*
Address
*/
func AddressCreate(address models.Address) int64 {
	db := getDBConn()

	log.Print("# Creating Address")

	stmt, err := db.Prepare("INSERT INTO address(address_country, address_zip, " +
		"address_state, address_city, address_line1, address_line2) " +
		"VALUES($1, $2, $3, $4, $5, $6) returning address_id")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	var address_id int64
	err = stmt.QueryRow(address.Country, address.Zipcode, address.State,
		address.City, address.Line1, address.Line2).Scan(&address_id)

	if err != nil {
		log.Panic(err)
	}

	return address_id
}

func AddressRead(address_id int64) models.Address {
	log.Print("# Reading Address")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT address_id, address_country, address_zip, " +
		"address_state, address_city, address_line1, address_line2 " +
		"FROM address WHERE address_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	var address = models.Address{}
	err = stmt.QueryRow(address_id).Scan(&address.AddressID, &address.Country,
		&address.Zipcode, &address.State, &address.City, &address.Line1, &address.Line2)

	if err == sql.ErrNoRows {
		return models.Address{}
	}

	if err != nil {
		log.Print("Error reading address data")
		log.Panic(err)
	}

	return address
}

func AddressUpdate(address models.Address) int64 {
	db := getDBConn()

	log.Print("# Updating Address")
	log.Printf("Address ID = %d", address.AddressID)

	stmt, err := db.Prepare("UPDATE address SET address_country = $1 " +
		" ,address_zip = $2 " +
		" ,address_state = $3 ,address_city = $4" +
		" ,address_line1 = $5 ,address_line2 = $6" +
		" WHERE address_id = $7")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Exec(address.Country, address.Zipcode, address.State,
		address.City, address.Line1, address.Line2, address.AddressID)

	if err != nil {
		log.Print("Address Update Failed")

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Panicf("Unexpected number of updates: %d", affectedCount)
	}

	return affectedCount
}

func AddressDelete(address_id int64) int64 {
	db := getDBConn()

	log.Print("# Deleting Address")
	log.Printf("Address ID = %d", address_id)

	stmt, err := db.Prepare("DELETE FROM address WHERE address_id = $1")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Print(err)
	}

	result, err := stmt.Exec(address_id)

	if err != nil {
		log.Print("Address Delete Failed")

		if isForeignKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		log.Panicf("Unexpected number of deletes: %d", affectedCount)
	}

	return affectedCount
}
