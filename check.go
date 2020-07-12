package check

import (
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strings"
)

// Formatter is an interface for logging failures, and then
// either abort or continue (Fatalf vs Errorf).
// The std lib testing.T and testing.B implements this
// interface implicitly.
type Formatter interface {
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

const errorMessage = "%s: Expected: %v. Actual: %v"

// Check validates so that expected and actual are equal.
// If they're not, an error will be logged and the test
// will be marked as failed, but continue to run.
func Check(t Formatter, expected, actual interface{}) {
	performValidation(expected, actual, t.Errorf)
}

// Assert validates so that expected and actual are equal.
// If they're not, an error will be logged and the test stopped.
func Assert(t Formatter, expected, actual interface{}) {
	performValidation(expected, actual, t.Fatalf)
}

func performValidation(expected, actual interface{}, log func(format string, args ...interface{})) {
	if !reflect.DeepEqual(expected, actual) {
		log(errorMessage, getLineInfo(), expected, actual)
	}
}

func getLineInfo() string {
	if _, path, line, ok := runtime.Caller(3); ok {
		fileName := getFileName(path)
		return fmt.Sprintf("%s:%d", fileName, line)
	}
	return "<Unknown Location>"
}

func getFileName(path string) string {
	split := strings.LastIndex(path, string(os.PathSeparator))
	if split == -1 || split == len(path) {
		return path
	}

	return path[split+1:]
}
