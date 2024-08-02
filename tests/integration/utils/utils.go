package utils

import (
	"fmt"
	"os"
	"reflect"
	"runtime/debug"
	"strings"
	"time"

	"asb-queue-emulator/tests/integration/base"
)

func AssertString(testSuite base.TestSuite, expected string, actual string, details ...interface{}) {
	if strings.Compare(expected, actual) != 0 {
		printErrorAndExit(testSuite, "Assertion failed. Expected: %s Actual: %s; Details: %v", expected, actual, details)
	}
}

// assertStringContains asserts that substr is within s
func AssertStringContains(testSuite base.TestSuite, s, substr string, details ...interface{}) {
	if !strings.Contains(s, substr) {
		printErrorAndExit(testSuite, "Assertion failed. '%s' should contain '%s'; Details: %v", s, substr, details)
	}
}

func AssertEqual(testSuite base.TestSuite, expected interface{}, actual interface{}, details ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		printErrorAndExit(testSuite, "Assertion failed. Expected: %v Actual: %v; Details: %v", expected, actual, details)
	}
}

func AssertNotEqual(testSuite base.TestSuite, expected interface{}, actual interface{}, details ...interface{}) {
	if reflect.DeepEqual(expected, actual) {
		printErrorAndExit(testSuite, "Assertion failed. Expected: not %v Actual: %v; Details: %v", expected, actual, details)
	}
}

func AssertTimeDifference(testSuite base.TestSuite, expected time.Time, actual time.Time, maxDifference time.Duration, details ...interface{}) {
	var difference time.Duration
	if expected.Before(actual) {
		difference = actual.Sub(expected)
	} else {
		difference = expected.Sub(actual)
	}
	if difference > maxDifference {
		printErrorAndExit(
			testSuite,
			"Assertion failed. Expected: %v Actual: %v Difference: %v Max Difference: %v; Details: %v",
			expected,
			actual,
			difference,
			maxDifference,
			details)
	}
}

func printErrorAndExit(testSuite base.TestSuite, format string, a ...interface{}) {
	fmt.Printf(format+"\n", a...)
	debug.PrintStack()
	testSuite.AfterSuite()
	os.Exit(1)
}

func IsNil(a interface{}) bool {
	if a == nil {
		return true
	}

	switch reflect.TypeOf(a).Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return reflect.ValueOf(a).IsNil()
	}

	return false
}

func AssertIsNil(testSuite base.TestSuite, a interface{}, details ...interface{}) {
	if !IsNil(a) {
		printErrorAndExit(testSuite, "Value %v is expected to be nil but it isn't. Details: %v", a, details)
	}
}

func AssertIsNotNil(testSuite base.TestSuite, a interface{}, details ...interface{}) {
	if IsNil(a) {
		printErrorAndExit(testSuite, "Value %v is expected to not be nil but it is. Details: %v", a, details)
	}
}
