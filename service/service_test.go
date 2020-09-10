package service

import (
    "testing"

    "github.com/stretchr/testify/assert"

    "mockexample/log"
)

// ... continued from log/logger_test.go
// Note: To run these tests, just run `cd service` and `go test` in the terminal
//       on the right.
//
// Whew! Almost done! This part is more incidental than anything, but is more of
// a way to show that we can use our mockable test logger with an ACTUAL *testing.T
// instance instead of just the mockTestingT we used in the test logger tests.

// This first test passes, notice how when you run `go test` the output from the
// call to `IsThingEmpty` isn't showing in the terminal output.
func TestServiceIsThingEmptyPassing(t *testing.T) {
    tl := log.NewMockableTestLogger(t)
    s := NewService(tl)

    assert.True(t, s.IsThingEmpty(""))
}

// This test fails on purpose. You'll see that this test's output includes the
// log message that comes from the `IsThingEmpty` call in the failed test output.
func TestServiceIsThingEmptyFailing(t *testing.T) {
    tl := log.NewMockableTestLogger(t)
    s := NewService(tl)

    assert.True(t, s.IsThingEmpty("this will fail"))
}

// Now to show that the TestLogger works the same as MockableTestLogger... But who
// really knows because we can't test it?! 😁
func TestServiceIsThingEmptyPassingNoMock(t *testing.T) {
    tl := log.NewTestLogger(t)
    s := NewService(tl)

    assert.True(t, s.IsThingEmpty(""))
}
func TestServiceIsThingEmptyFailingNoMock(t *testing.T) {
    tl := log.NewTestLogger(t)
    s := NewService(tl)

    assert.True(t, s.IsThingEmpty("this will fail"))
}