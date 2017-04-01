package dbhandler

import (
	"testing"
)

func TestCredentialCreate(t *testing.T) {
	var tests = []credential_struct {
		{1, "test@ksu.edu", "test"},
		{2, "test1@ksu.edu", "test1"},
	}

	for _, c := range tests {
		got := CredentialCreate(c.emailaddress, c.password_hash)

		if (got != c.credential_id) {
			t.Errorf("Inserted %s with expected ID %d, but got %d",c.emailaddress, c.credential_id, got)
		}
	}
}

func TestCredentialRead(t *testing.T) {
	var tests = [] credential_struct {
		{1, "test@ksu.edu", "test"},
		{2, "test1@ksu.edu", "test1"},
	}

	for _, c := range tests {
		got := CredentialRead(c.emailaddress)

		if (got.password_hash != c.password_hash) {
			t.Errorf("Queried emailaddress = %s with expected password %s, but got %s",c.emailaddress, c.password_hash, got.password_hash)
		}
	}
}

func TestCredentialUpdate(t *testing.T) {
	var tests = []credential_struct {
		{1, "test1@ksu.edu", "test2"},
	}

	for _, c := range tests {
		_ = CredentialUpdate(c.emailaddress, c.password_hash)

		got := CredentialRead(c.emailaddress)

		if (got.password_hash != c.password_hash) {
			t.Errorf("Updated emailadress = %s with new password %s, but got %s",c.emailaddress, c.password_hash, got)
		}
	}
}

func TestCredentialDelete(t *testing.T) {
    var tests = []credential_struct {
        {1, "test1@ksu.edu", "test2"},
    }

	for _, c := range tests {
		_ = CredentialDelete(c.emailaddress)

		got := CredentialRead(c.emailaddress)

		if (got.emailaddress == c.emailaddress) {
			t.Errorf("Deleted emailaddress = %s but still got %s",c.emailaddress, got)
		}
	}
}

func TestAddressCreate(t *testing.T) {
    var addresses = []address_struct {
        {1, "USA", "66502", "KS", "Manhattan", "2100 Poytz Avenue", ""},
    }

    for _, c := range addresses {
        got := AddressCreate(c)

        if (got != c.address_id) {
            t.Errorf("Created address with expected ID %d, bit got %d", c.address_id, got)
        }
    }
}

func TestAddressRead(t *testing.T) {
    var addresses = []address_struct {
        {1, "USA", "66502", "KS", "Manhattan", "2100 Poytz Avenue", ""},
    }

    for _, c := range addresses {
        got := AddressRead(c.address_id)

        if (got.address_id != c.address_id) {
            t.Errorf("Created address with expected ID %d, bit got %d", c.address_id, got.address_id)
        }
    }
}

func TestAddressUpdate(t *testing.T) {
    var addresses = []address_struct {
        {1, "USA", "66502", "KS", "Manhattan", "2100", "Poytz Avenue"},
    }

    for _, c := range addresses {
        got := AddressUpdate(c)

        if (got.address_id != c.address_id) {
            t.Errorf("Created address with expected ID %d, bit got %d", c.address_id, got.address_id)
        }
    }
}

func TestAddressDelete(t *testing.T) {
    var addresses = []address_struct {
        {1, "USA", "66502", "KS", "Manhattan", "2100", "Poytz Avenue"},
    }

    for _, c := range addresses {
        got := AddressDelete(c.address_id)

        if (got != c.address_id) {
            t.Errorf("Deleted address with expected ID %d, bit got %d", c.address_id, got)
        }
    }
}