package dbhandler

import (
	"testing"
)

/**
Credential
*/
func TestCredentialCreate(t *testing.T) {
	var tests = []credential_struct{
		{1, "test@ksu.edu", "test"},
		{2, "test1@ksu.edu", "test1"},
	}

	for _, c := range tests {
		got := CredentialCreate(c.emailaddress, c.password_hash)

		if got != c.credential_id {
			t.Errorf("Inserted %s with expected ID %d, but got %d", c.emailaddress, c.credential_id, got)
		}
	}
}

func TestCredentialRead(t *testing.T) {
	var tests = []credential_struct{
		{1, "test@ksu.edu", "test"},
		{2, "test1@ksu.edu", "test1"},
	}

	for _, c := range tests {
		got := CredentialRead(c.emailaddress)

		if got.password_hash != c.password_hash {
			t.Errorf("Queried emailaddress = %s with expected password %s, but got %s", c.emailaddress, c.password_hash, got.password_hash)
		}
	}
}

func TestCredentialUpdate(t *testing.T) {
	var tests = []credential_struct{
		{1, "test1@ksu.edu", "test2"},
	}

	for _, c := range tests {
		_ = CredentialUpdate(c.emailaddress, c.password_hash)

		got := CredentialRead(c.emailaddress)

		if got.password_hash != c.password_hash {
			t.Errorf("Updated emailadress = %s with new password %s, but got %s", c.emailaddress, c.password_hash, got)
		}
	}
}

func TestCredentialDelete(t *testing.T) {
	var tests = []credential_struct{
		{1, "test1@ksu.edu", "test2"},
	}

	for _, c := range tests {
		_ = CredentialDelete(c.emailaddress)

		got := CredentialRead(c.emailaddress)

		if got.emailaddress == c.emailaddress {
			t.Errorf("Deleted emailaddress = %s but still got %s", c.emailaddress, got)
		}
	}
}

/**
Address
*/
func TestAddressCreate(t *testing.T) {
	var addresses = []address_struct{
		{1, "USA", "66502", "KS", "Manhattan", "2100 Poytz Avenue", ""},
		{2, "USA", "67601", "KS", "Hays", "600", "Park Street"},
	}

	for _, c := range addresses {
		got := AddressCreate(c)

		if got != c.address_id {
			t.Errorf("Created address with expected ID %d, but got %d", c.address_id, got)
		}
	}
}

func TestAddressRead(t *testing.T) {
	var addresses = []address_struct{
		{1, "USA", "66502", "KS", "Manhattan", "2100 Poytz Avenue", ""},
	}

	for _, c := range addresses {
		got := AddressRead(c.address_id)

		if got.line1 != c.line1 {
			t.Errorf("Tried reading address with ID %d and expected line1 %s but got %s", c.address_id, c.line1, got.line1)
		}
	}
}

func TestAddressUpdate(t *testing.T) {
	var addresses = []address_struct{
		{1, "USA", "66502", "KS", "Manhattan", "2100", "Poytz Avenue"},
	}

	for _, c := range addresses {
		_ = AddressUpdate(c)

		got := AddressRead(c.address_id)

		if got.line1 != c.line1 {
			t.Errorf("Updated address ID %d with expected line1 as %s, but got %s", c.address_id, c.line1, got.line1)
		}
	}
}

func TestAddressDelete(t *testing.T) {
	var addresses = []address_struct{
		{1, "USA", "66502", "KS", "Manhattan", "2100", "Poytz Avenue"},
	}

	for _, c := range addresses {
		_ = AddressDelete(c.address_id)

		got := AddressRead(c.address_id)

		if got.address_id == c.address_id {
			t.Errorf("Deleted address with expected ID %d, but got %d", c.address_id, got.address_id)
		}
	}
}

/**
School
*/

func TestSchoolCreate(t *testing.T) {
	var schools = []school_struct{
		{1, "Manhattan High School", 2},
		{2, "Kansas Academy of Mathematics and Science", 2},
		{3, "De Soto High Schoo", 2},
		{4, "Andover Central High School", 2},
	}

	for _, c := range schools {
		got := SchoolCreate(c)

		if got != c.school_id {
			t.Errorf("Created School with expected ID %d, but got %d", c.school_id, got)
		}
	}
}

func TestSchoolRead(t *testing.T) {
	var schools = []school_struct{
		{1, "Manhattan High School", 1},
		{2, "Kansas Academy of Mathematics and Science", 1},
		{3, "De Soto High Schoo", 1},
		{4, "Andover Central High School", 1},
	}

	for _, c := range schools {
		got := SchoolRead(c.school_id)

		if got.school_name != c.school_name {
			t.Errorf("Queried School ID %d with expected name of %s, but got %s", c.school_id, c.school_name, got.school_name)
		}
	}
}

func TestSchoolUpdate(t *testing.T) {
	// create a new address
	var address address_struct = address_struct{
		3, "USA", "66502", "KS", "Manhattan", "2100 Poytz Avenue", "",
	}
	// save the address_id
	var new_address_id int64
	new_address_id = AddressCreate(address)

	var schools = []school_struct{
		{2, "Kansas Academy of Mathematics and Science", new_address_id},
	}

	for _, c := range schools {
		_ = SchoolUpdate(c)

		got := SchoolRead(c.school_id)

		if got.address_id != c.address_id {
			t.Errorf("Updated School ID %d with address ID of %d, but got %d", c.school_id, c.address_id, got.address_id)
		}
	}
}

func TestSchoolDelete(t *testing.T) {
	var schools = []school_struct{
		{4, "Andover Central High School", 1},
	}

	for _, c := range schools {
		_ = SchoolDelete(c.school_id)

		got := SchoolRead(c.school_id)

		if got.school_name == c.school_name {
			t.Errorf("Deleted School ID %d but got %d", c.school_id, got.school_id)
		}
	}
}
