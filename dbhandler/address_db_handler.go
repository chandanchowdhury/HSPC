package dbhandler

import (
	"database/sql"
	"github.com/chandanchowdhury/HSPC/models"
	"log"
)

//TODO: More data validation

/**
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
		log.Panic(err)
	}

	var address_id int64
	err = stmt.QueryRow(address.Country, address.Zipcode, address.State,
		address.City, address.Line1, address.Line2).Scan(&address_id)

	if err != nil {
		//if the address already exists, instead of returning error
		// return the existing address
		if isDuplicateKeyError(err) {
			addr := FindAddress(address.Zipcode, address.Line1)
			return addr.AddressID
		}

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
		log.Panic(err)
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
		log.Panic(err)
	}

	result, err := stmt.Exec(address.Country, address.Zipcode, address.State,
		address.City, address.Line1, address.Line2, address.AddressID)

	if err != nil {
		log.Print("Address Update Failed")

		if isForeignKeyError(err) {
			return -1
		}

		if isDuplicateKeyError(err) {
			return -2
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		// if no records updated, just inform the caller
		if affectedCount == 0 {
			return 0
		}

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
		log.Panic(err)
	}

	result, err := stmt.Exec(address_id)

	if err != nil {
		log.Print("Address Delete Failed")

		if isForeignKeyError(err) {
			return -1
		}

		log.Panic(err)
	}

	affectedCount, err := result.RowsAffected()

	if affectedCount != 1 {
		// if no records deleted, well nothing to do then, just
		// inform the caller
		if affectedCount == 0 {
			return 0
		}

		log.Panicf("Unexpected number of deletes: %d", affectedCount)
	}

	return affectedCount
}

/**
Given zipcode and address line 1 we should be able to find an address.
We assume that unlike houses/apartments there cannot be two schools at
same zip code (state and city) and same address line1.
*/
func FindAddress(address_zip *string, address_line1 *string) models.Address {
	log.Print("# Reading Address")

	db := getDBConn()
	stmt, err := db.Prepare("SELECT address_id, address_country, address_zip, " +
		"address_state, address_city, address_line1, address_line2 FROM address " +
		"WHERE address_zip = $1 AND address_line1 = $2")
	defer stmt.Close()

	if err != nil {
		log.Print("Error creating prepared statement")
		log.Panic(err)
	}

	var address = models.Address{}
	err = stmt.QueryRow(address_zip, address_line1).Scan(&address.AddressID, &address.Country,
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
