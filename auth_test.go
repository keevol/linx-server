package main

import (
	"testing"
)

func TestCheckAuth(t *testing.T) {
	authKeys := []string{
		"vhvZ/PT1jeTbTAJ8JdoxddqFtebSxdVb0vwPlYO+4HM=",
		"vFpNprT9wbHgwAubpvRxYCCpA2FQMAK6hFqPvAGrdZo=",
	}

	if r, err := checkAuth(authKeys, ""); err != nil && r {
		t.Fatal("Authorization passed for empty key")
	}

	if r, err := checkAuth(authKeys, "thisisnotvalid"); err != nil && r {
		t.Fatal("Authorization passed for invalid key")
	}

	if r, err := checkAuth(authKeys, "haPVipRnGJ0QovA9nyqK"); err != nil && !r {
		t.Fatal("Authorization failed for valid key")
	}
}
