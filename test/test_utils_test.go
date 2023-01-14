package test

import (
	"testing"

	"github.com/case-management-suite/testutil"
)

func TestCounterSeqUtil(t *testing.T) {
	n := 10
	reduce, current, done := testutil.CounterSeqUtil(n, func(int) {})

	total := (n * (n + 1)) / 2

	for i := 1; i <= 10; i++ {
		c, err := reduce(i)
		testutil.AssertNilError(err, t)
		total -= i
		testutil.AssertEq(total, c, t)
		testutil.AssertEq(total, current(), t)
	}
	testutil.AssertEq(0, current(), t)
	testutil.AssertTrue(done(), t)
}
