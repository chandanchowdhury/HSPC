package dbhandler

import (
	"testing"
)

func TestEmailAddressInsert(t *testing.T) {
	var tests = []struct {
		emailaddress string
		emailaddress_id int
	} {
		{"test@ksu.edu", 1},
		{"test1@ksu.edu", 2},
	}

	for _, c := range tests {
		got := EmailAddressCreate(c.emailaddress)

		if (got != c.emailaddress_id) {
			t.Errorf("Inserted %s with expected ID %i, but got %i",c.emailaddress, c.emailaddress_id, got)
		}
	}
}

func TestEmailAddressGet(t *testing.T) {
	var tests = []struct {
		emailaddress string
		emailaddress_id int
	} {
		{"test@ksu.edu", 1},
		{"test1@ksu.edu", 2},
	}

	for _, c := range tests {
		got := EmailAddressRead(c.emailaddress_id)

		if (got != c.emailaddress) {
			t.Errorf("Queried ID = %d with expected emailaddress %s, but got %s",c.emailaddress_id, c.emailaddress, got)
		}
	}
}

func TestEmailAddressUpdate(t *testing.T) {
	var tests = []struct {
		emailaddress string
		emailaddress_id int
	} {
		{"test2@ksu.edu", 2},
	}

	for _, c := range tests {
		_ = EmailAddressUpdate(c.emailaddress_id, c.emailaddress)

		got := EmailAddressRead(c.emailaddress_id)

		if (got != c.emailaddress) {
			t.Errorf("Updated ID = %d with new emailaddress %s, but got %s",c.emailaddress_id, c.emailaddress, got)
		}
	}
}

func TestEmailAddressDelete(t *testing.T) {
	var tests = []struct {
		emailaddress string
		emailaddress_id int
	} {
		{"test2@ksu.edu", 2},
	}

	for _, c := range tests {
		_ = EmailAddressDelete(c.emailaddress_id)

		got := EmailAddressRead(c.emailaddress_id)

		if (got == c.emailaddress) {
			t.Errorf("Deleted ID = %d with new emailaddress %s, but still got %s",c.emailaddress_id, c.emailaddress, got)
		}
	}
}