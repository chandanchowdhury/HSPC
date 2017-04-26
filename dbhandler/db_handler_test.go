package dbhandler

import (
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/chandanchowdhury/HSPC/dbhandler"
	"github.com/go-openapi/strfmt"
	"testing"
)

/**
Credential
*/
var false_var = false

//var true_var = true;

func TestCredentialCreate(t *testing.T) {
	var email1 strfmt.Email = strfmt.Email("test@ksu.edu")
	var email2 strfmt.Email = strfmt.Email("test1@ksu.edu")
	var password1 strfmt.Password = strfmt.Password("test")
	var password2 strfmt.Password = strfmt.Password("test1")
	var tests = []models.Credential{
		{1, &email1, &password1, &false_var},
		{2, &email2, &password2, &false_var},
	}

	for _, c := range tests {
		got := dbhandler.CredentialCreate(c)

		if got != c.CredentialID {
			t.Errorf("Inserted %s with expected ID %d, but got %d", c.Emailaddress, c.CredentialID, got)
		}
	}
}

func TestCredentialRead(t *testing.T) {
	var email1 strfmt.Email = strfmt.Email("test@ksu.edu")
	var email2 strfmt.Email = strfmt.Email("test1@ksu.edu")
	var password1 strfmt.Password = strfmt.Password("test")
	var password2 strfmt.Password = strfmt.Password("test1")

	var tests = []models.Credential{
		{1, &email1, &password1, &false_var},
		{2, &email2, &password2, &false_var},
	}

	for _, c := range tests {
		got := CredentialRead(c.Emailaddress.String())

		if got.Password.String() != c.Password.String() {
			t.Errorf("Queried emailaddress = %s with expected password %s, but got %s", c.Emailaddress, c.Password, got.Password)
		}
	}
}

func TestCredentialUpdate(t *testing.T) {
	var email1 strfmt.Email = strfmt.Email("test@ksu.edu")
	var password2 strfmt.Password = strfmt.Password("test2")
	var tests = []models.Credential{
		{1, &email1, &password2, &false_var},
	}

	for _, c := range tests {
		_ = CredentialUpdate(c.Emailaddress.String(), c.Password.String())

		got := CredentialRead(c.Emailaddress.String())

		if got.Password.String() != c.Password.String() {
			t.Errorf("Updated emailadress = %s with new password %s, but got %s", c.Emailaddress, c.Password, got.Password)
		}
	}
}

func TestCredentialDelete(t *testing.T) {
	var email1 strfmt.Email = strfmt.Email("test@ksu.edu")
	var password2 strfmt.Password = strfmt.Password("test2")

	var tests = []models.Credential{
		{1, &email1, &password2, &false_var},
	}

	for _, c := range tests {
		_ = CredentialDelete(c.Emailaddress.String())

		got := CredentialRead(c.Emailaddress.String())

		if got.Emailaddress == c.Emailaddress {
			t.Errorf("Deleted emailaddress = %s but still got %s", c.Emailaddress.String(), got.Emailaddress.String())
		}
	}
}

/**
Address
*/
func TestAddressCreate(t *testing.T) {
	var country = "USA"
	var zipcodes = []string{"66502", "67601"}
	var state = []string{"KS"}
	var city = []string{"Manhattan", "Hays"}
	var line1 = []string{"2100 Poytz Avenue", "600"}
	var line2 = []string{"", "Park Street"}

	var addresses = []models.Address{
		{1, &country, &zipcodes[0], &state[0], city[0], &line1[0], &line2[0]},
		{2, &country, &zipcodes[1], &state[0], city[1], &line1[1], &line2[1]},
	}

	for _, c := range addresses {
		got := AddressCreate(c)

		if got != c.AddressID {
			t.Errorf("Created address with expected ID %d, but got %d", c.AddressID, got)
		}
	}
}

func TestAddressRead(t *testing.T) {
	var country = "USA"
	var zipcodes = []string{"66502", "67601"}
	var state = []string{"KS"}
	var city = []string{"Manhattan", "Hays"}
	var line1 = []string{"2100 Poytz Avenue", "600"}
	var line2 = []string{"", "Park Street"}

	var addresses = []models.Address{
		{1, &country, &zipcodes[0], &state[0], city[0], &line1[0], &line2[0]},
	}

	for _, c := range addresses {
		got := AddressRead(c.AddressID)

		if *got.Line1 != *c.Line1 {
			t.Errorf("Tried reading address with ID %d and expected Line1 %s but got %s", c.AddressID, *c.Line1, *got.Line1)
		}
	}
}

func TestAddressUpdate(t *testing.T) {
	var country = "USA"
	var zipcodes = []string{"66502", "67601"}
	var state = []string{"KS"}
	var city = []string{"Manhattan", "Hays"}
	var line1 = []string{"Poytz Avenue", "600"}
	var line2 = []string{"2100", "Park Street"}

	var addresses = []models.Address{
		{1, &country, &zipcodes[0], &state[0], city[0], &line1[0], &line2[0]},
	}

	for _, c := range addresses {
		_ = AddressUpdate(c)

		got := AddressRead(c.AddressID)

		if *got.Line1 != *c.Line1 {
			t.Errorf("Updated address ID %d with expected Line1 as %s, but got %s", c.AddressID, *c.Line1, *got.Line1)
		}
	}
}

func TestAddressDelete(t *testing.T) {
	var country = "USA"
	var zipcodes = []string{"66502", "67601"}
	var state = []string{"KS"}
	var city = []string{"Manhattan", "Hays"}
	var line1 = []string{"Poytz Avenue", "600"}
	var line2 = []string{"2100", "Park Street"}

	var addresses = []models.Address{
		{1, &country, &zipcodes[0], &state[0], city[0], &line1[0], &line2[0]},
	}

	for _, c := range addresses {
		_ = AddressDelete(c.AddressID)

		got := AddressRead(c.AddressID)

		if got.AddressID == c.AddressID {
			t.Errorf("Deleted address with expected ID %d, but got %d", c.AddressID, got.AddressID)
		}
	}
}

/**
School
*/

func TestSchoolCreate(t *testing.T) {
	var school_name = []string{
		"Manhattan High School", "Kansas Academy of Mathematics and Science",
		"De Soto High Schoo", "Andover Central High School"}

	var address_id = []int64{2}

	var schools = []models.School{
		{1, &school_name[0], &address_id[0], 0, false},
		{2, &school_name[1], &address_id[0], 0, false},
		{3, &school_name[2], &address_id[0], 0, false},
		{4, &school_name[3], &address_id[0], 0, false},
	}

	for _, c := range schools {
		got := SchoolCreate(c)

		if got != c.SchoolID {
			t.Errorf("Created School with expected ID %d, but got %d", c.SchoolID, got)
		}
	}
}

func TestSchoolRead(t *testing.T) {
	var school_name = []string{
		"Manhattan High School", "Kansas Academy of Mathematics and Science",
		"De Soto High Schoo", "Andover Central High School"}

	var address_id = []int64{2}

	var schools = []models.School{
		{1, &school_name[0], &address_id[0], 0, false},
		{2, &school_name[1], &address_id[0], 0, false},
		{3, &school_name[2], &address_id[0], 0, false},
		{4, &school_name[3], &address_id[0], 0, false},
	}

	for _, c := range schools {
		got := SchoolRead(c.SchoolID)

		if *got.SchoolName != *c.SchoolName {
			t.Errorf("Queried School ID %d with expected name of %s, but got %s", c.SchoolID, c.SchoolName, got.SchoolName)
		}
	}
}

func TestSchoolUpdate(t *testing.T) {
	var school_name = []string{
		"Manhattan High School", "Kansas Academy of Mathematics and Science",
		"De Soto High Schoo", "Andover Central High School"}

	var address_id = []int64{1}

	var schools = []models.School{
		{2, &school_name[1], &address_id[0], 0, false},
	}

	for _, c := range schools {
		_ = SchoolUpdate(c)

		got := SchoolRead(c.SchoolID)

		if got.AddressID != c.AddressID {
			t.Errorf("Updated School ID %d with address ID of %d, but got %d", c.SchoolID, c.AddressID, got.AddressID)
		}
	}
}

func TestSchoolDelete(t *testing.T) {
	var school_name = []string{"Andover Central High School"}

	var address_id = []int64{1}

	var schools = []models.School{
		{4, &school_name[0], &address_id[0], 0, false},
	}

	for _, c := range schools {
		_ = SchoolDelete(c.SchoolID)

		got := SchoolRead(c.SchoolID)

		if *got.SchoolName == *c.SchoolName {
			t.Errorf("Deleted School ID %d but got %d", c.SchoolID, got.SchoolID)
		}
	}
}
