package testutil

import (
	"testing"

	"github.com/case-management-suite/common/server"
	"golang.org/x/exp/constraints"
)

var TU = server.NewTestServerUtils()

func AssertEq[T constraints.Ordered](a T, b T, t *testing.T) {
	if a != b {
		TU.Logger.Error().Interface("actual", b).Interface("expected", a).Msg("Assertion failure")
		t.FailNow()
	}
}

func AssertTrue(value bool, t *testing.T) {
	if !value {
		TU.Logger.Error().Bool("actual", value).Msg("Assertion failure")
		t.FailNow()
	}
}

func AssertNonNil(value interface{}, t *testing.T) {
	if value == nil {
		TU.Logger.Error().Interface("actual", value).Msg("Assertion failure")
		t.FailNow()
	}
}

func AssertNotNilError(err error, t *testing.T) {
	if err == nil {
		TU.Logger.Error().Err(err).Msg("Expected error, but was nil")
		t.FailNow()
	}
}

func AssertNilError(err error, t *testing.T) {
	if err != nil {
		TU.Logger.Error().Err(err).Msg("Unexpected error")
		t.FailNow()
	}
}
