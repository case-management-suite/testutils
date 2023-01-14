package testutil

import (
	"errors"
	"flag"
	"time"
)

type TestConfig struct {
	RunIntegration bool
}

func GetTestConfig() TestConfig {
	config := TestConfig{}

	flag.BoolVar(&config.RunIntegration, "integration", false, "Run integration tests")
	return config
}

type ReduceFn = func(int) (int, error)
type CurrentFn = func() int
type DoneFn = func() bool

func CounterSeqUtil(n int, fn func(int)) (ReduceFn, CurrentFn, DoneFn) {
	total := (n * (n + 1)) / 2
	for i := 1; i <= n; i++ {
		fn(i)
	}
	return func(i int) (int, error) {
			total -= i
			if total < 0 {
				return total, errors.New("count went down by more than expected")
			}
			return total, nil
		},
		func() int {
			return total
		},
		func() bool {
			return total == 0
		}
}

func CounterUtil(n int, fn func(int)) (ReduceFn, CurrentFn, DoneFn) {
	for i := 1; i <= n; i++ {
		fn(i)
	}
	total := n
	return func(i int) (int, error) {
			total -= i
			if total < 0 {
				return total, errors.New("count went down by more than expected")
			}
			return total, nil
		},
		func() int {
			return total
		},
		func() bool {
			return total == 0
		}
}

func WithRetry[T any](
	retries int,
	cond func(T, error) bool,
	fn func() (T, error),
) (T, error) {
	var r T
	var e error
	for retries > 0 {
		retries--
		r, e := fn()
		if e == nil && cond(r, e) {
			return r, nil
		}
		time.Sleep(time.Second)
	}
	return r, e
}
