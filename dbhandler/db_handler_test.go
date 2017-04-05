package dbhandler

import (
	"testing"
	"github.com/chandanchowdhury/HSPC/models"
	"github.com/go-openapi/strfmt"
)

/**
Credential
*/
func TestCredentialCreate(t *testing.T) {
	var email1 strfmt.Email = strfmt.Email("test@ksu.edu")
	var email2 strfmt.Email = strfmt.Email("test1@ksu.edu")
	var tests = []models.Credential {
		{1, &email1, "test"},
		{2, &email2, "test1"},
	}

	for _, c := range tests {
		got := CredentialCreate(c)

		if got != c.CredentialID {
			t.Errorf("Inserted %s with expected ID %d, but got %d", c.Emailaddress, c.CredentialID, got)
		}
	}
}

func TestCredentialRead(t *testing.T) {
	var email1 strfmt.Email = strfmt.Email("test@ksu.edu")
	var email2 strfmt.Email = strfmt.Email("test1@ksu.edu")
	var tests = []models.Credential {
		{1, &email1, "test"},
		{2, &email2, "test1"},
	}

	for _, c := range tests {
		got := CredentialRead(c.Emailaddress.String())

		if got.Password != c.Password {
			t.Errorf("Queried emailaddress = %s with expected password %s, but got %s", c.Emailaddress, c.Password, got.Password)
		}
	}
}

func TestCredentialUpdate(t *testing.T) {
	var email1 strfmt.Email = strfmt.Email("test@ksu.edu")
	var tests = []models.Credential {
		{1, &email1, "test2"},
	}

	for _, c := range tests {
		_ = CredentialUpdate(c.Emailaddress.String(), c.Password.String())

		got := CredentialRead(c.Emailaddress.String())

		if got.Password != c.Password {
			t.Errorf("Updated emailadress = %s with new password %s, but got %s", c.Emailaddress, c.Password, got.Password)
		}
	}
}

func TestCredentialDelete(t *testing.T) {
	var email1 strfmt.Email = strfmt.Email("test@ksu.edu")
	var tests = []models.Credential{
		{1, &email1, "test2"},
	}

	for _, c := range tests {
		_ = CredentialDelete(c.Emailaddress.String())

		got := CredentialRead(c.Emailaddress.String())

		if got.Emailaddress == c.Emailaddress {
			t.Errorf("Deleted emailaddress = %s but still got %s", c.Emailaddress.String(), got.Emailaddress.String())
		}
	}
}
