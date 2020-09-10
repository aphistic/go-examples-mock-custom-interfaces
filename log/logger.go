package log

import (
    "fmt"
    "testing"
)

type Logger interface {
    Debug(msg string)
    Info(msg string)
    Error(msg string)
}

// Our "standard" logger that logs to stdout
type StdoutLogger struct {}

// Let the compiler tell us if StdoutLogger implements Logger or not
var _ Logger = &TestLogger{}

func NewStdoutLogger() *StdoutLogger {
    return &StdoutLogger{}
}

func (sl *StdoutLogger) Debug(msg string) {
    fmt.Printf("[DEBUG] " + msg + "\n")
}
func (sl *StdoutLogger) Info(msg string) {
    fmt.Printf("[INFO] " + msg + "\n")
}
func (sl *StdoutLogger) Error(msg string) {
    fmt.Printf("[ERROR] " + msg + "\n")
}

// When we're testing our application we don't want it logging things
// to stdout and instead to each individual test we write. In order to do this
// we need something that fulfills the Logger interface, but takes a *testing.T
// so we can use the testing.T.Log method to log output. This is great/fine for actual
// usage but how do we know our test logger is actually working correctly? We need to
// test the test logger! However, because it only takes a *testing.T we can't actually
// mock out the testing.T instance to see what it's actually logging. For this reason
// we'll want to create a test logger we can mock instead. This would usually replace
// the testLogger implementation below but I want to keep both of them for example
// purposes.
type TestLogger struct {
    t *testing.T
}

// Let the compiler tell us if TestLogger implements Logger or not
var _ Logger = &TestLogger{}

func NewTestLogger(t *testing.T) *TestLogger {
    return &TestLogger{
        t: t,
    }
}

func (tl *TestLogger) Debug(msg string) {
    tl.t.Logf("[DEBUG(test)] " + msg)
}
func (tl *TestLogger) Info(msg string) {
    tl.t.Logf("[INFO(test)] " + msg)
}
func (tl *TestLogger) Error(msg string) {
    tl.t.Logf("[ERROR(test)] " + msg)
}

// Here's where we start the code we'd actually use to create a mockable test logger.

// First, create an interface that fits what we care about from *testing.T. Since Go has
// implicit interface implementations we can use this anywhere we'd want to use a *testing.T
// even though *testing.T doesn't have an "official" interface. One thing to remember here is
// that we don't need to define every method that's on *testing.T, only the ones we actually
// use or care about. In this case it's only the Logf method.
type TestingT interface {
    Logf(format string, v ...interface{})
}

// Now create our struct that uses our TestingT interface. For the next steps take a look at
// the logger_test.go file to see how we actually test this struct.
type MockableTestLogger struct {
    t TestingT
}

// Let the compiler tell us if MockableTestLogger implements Logger or not
var _ Logger = &TestLogger{}

func NewMockableTestLogger(t TestingT) *MockableTestLogger {
    return &MockableTestLogger{
        t: t,
    }
}

func (tl *MockableTestLogger) Debug(msg string) {
    tl.t.Logf("[DEBUG(test)] " + msg)
}
func (tl *MockableTestLogger) Info(msg string) {
    tl.t.Logf("[INFO(test)] " + msg)
}
func (tl *MockableTestLogger) Error(msg string) {
    tl.t.Logf("[ERROR(test)] " + msg)
}