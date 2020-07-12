package check

import (
	"fmt"
	"runtime"
	"testing"
)

// testingMock is used so that we can validate the result
// of calling Check and Assert.
type testingMock struct {
	result string
	called bool
}

func (t *testingMock) Errorf(format string, args ...interface{}) {
	t.called = true
	t.result = fmt.Sprintf(format, args...)
}

func (t *testingMock) Fatalf(format string, args ...interface{}) {
	t.called = true
	t.result = fmt.Sprintf(format, args...)
}

func TestCheck(t *testing.T) {
	mock := &testingMock{}
	Check(mock, 1, 1)

	// Irony of not using Check is not lost on me
	if mock.called {
		t.Error("Didn't expect testing framework to be called.")
	}

	Check(mock, 1, 2)
	_, _, num, _ := runtime.Caller(0)
	expected := fmt.Sprintf("check_test.go:%d: Expected: 1. Actual: 2", num-1)
	if mock.result != expected {
		t.Errorf("Expected: %s\nActual:%s\n", expected, mock.result)
	}
}
