package commands

import (
	"testing"
)

func assertError(t testing.TB, err error, wantError bool) {
	if (err != nil) != wantError {
		t.Errorf("wantError? = %t, got: %v", wantError, err)
	}
}
