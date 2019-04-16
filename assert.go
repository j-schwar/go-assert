// Package assert provides an assortment of assertions to be used in conjunction
// with golang's testing package to streamline tests.
package assert

import (
	"testing"
)

// IsError fails the test if a given error is nil.
func IsError(t *testing.T, err error) {
	if err == nil {
		t.Error(msgIsError)
	}
}

// FatalIsError fails the test immediately if a given error is nil.
func FatalIsError(t *testing.T, err error) {
	if err == nil {
		t.Fatal(msgIsError)
	}
}

// NotError fails the test if a given error is not nil.
func NotError(t *testing.T, err error) {
	if err != nil {
		t.Errorf(formatIsNotError, err)
	}
}

// FatalNotError fails the test immediately if a given error is not nil.
func FatalNotError(t *testing.T, err error) {
	if err != nil {
		t.Fatalf(formatIsNotError, err)
	}
}

// CondError checks that a given error was expected or not.
//
// If expected is true and err is nil, the test will fail. Conversely, if
// expected is false and err is not nil, the test will also fail.
func CondError(t *testing.T, expected bool, err error) {
	if expected && err == nil {
		t.Error(msgIsError)
	} else if !expected && err != nil {
		t.Errorf(formatIsNotError, err)
	}
}

// FatalCondError checks that a given error was expected or not.
//
// If expected is true and err is nil, the test will fail immediately.
// Conversely, if expected is false and err is not nil, the test will also fail
// immediately.
func FatalCondError(t *testing.T, expected bool, err error) {
	if expected && err == nil {
		t.Fatal(msgIsError)
	} else if !expected && err != nil {
		t.Fatalf(formatIsNotError, err)
	}
}
