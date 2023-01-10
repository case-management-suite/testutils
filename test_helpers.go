package testutil

import (
	"testing"

	"github.com/rs/zerolog/log"
	"golang.org/x/exp/constraints"
)

func AssertEq[T constraints.Ordered](a T, b T, t *testing.T) {
	if a != b {
		log.Fatal().Interface("actual", b).Interface("expected", a).Msg("Assertion failure")
		t.FailNow()
	}
}
