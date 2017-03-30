package dbhandler

import (
	"testing"
)

func TestCredentialInsert(t *testing.T) {
	var tests = []struct {
		emailaddress string
		password string
		credential_id int
	} {
		{"test@ksu.edu", "test", 1},
		{"test1@ksu.edu", "test1", 2},
	}

	for _, c := range tests {
		got := CredentialCreate(c.emailaddress, c.password)

		if (got != c.credential_id) {
			t.Errorf("Inserted %s with expected ID %i, but got %i",c.emailaddress, c.emailaddress, got)
		}
	}
}

func TestCredentialGet(t *testing.T) {
	var tests = []struct {
		emailaddress string
		password string
	} {
		{"test@ksu.edu", "test"},
		{"test1@ksu.edu", "test1"},
	}

	for _, c := range tests {
		got := CredentialRead(c.emailaddress)

		if (got != c.password) {
			t.Errorf("Queried emailaddress = %s with expected password %s, but got %s",c.emailaddress, c.password, got)
		}
	}
}

func TestCredentialUpdate(t *testing.T) {
	var tests = []struct {
		emailaddress string
		password string
	} {
		{"test1@ksu.edu", "test2"},
	}

	for _, c := range tests {
		_ = CredentialUpdate(c.emailaddress, c.password)

		got := CredentialRead(c.emailaddress)

		if (got != c.password) {
			t.Errorf("Updated emailadress = %s with new password %s, but got %s",c.emailaddress, c.password, got)
		}
	}
}

func TestCredentialDelete(t *testing.T) {
	var tests = []struct {
		emailaddress string
	} {
		{"test1@ksu.edu"},
	}

	for _, c := range tests {
		_ = CredentialDelete(c.emailaddress)

		got := CredentialRead(c.emailaddress)

		if (got == c.emailaddress) {
			t.Errorf("Deleted emailaddress = %s but still got %s",c.emailaddress, got)
		}
	}
}