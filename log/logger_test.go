package log

import (
    "fmt"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

// ...continued from logger.go
// Note: If you'd just like to run these tests, type `cd log` and then `go test` into
//       the terminal on the right.
//
// Now that both our un-mockable test logger and mockable test loggers are
// created, it's time to test them! But... the mockable one is really the only
// one we can test. :)

// The first step in testing it is to create our "mock" *testing.T implementation
// we can use to check that the test logger is logging correctly:
type mockTestingT struct {
    // messages contains all the calls to Log
    messages []string
}

func (mtt *mockTestingT) Logf(format string, v ...interface{}) {
    // Add the log message to our messages
    mtt.messages = append(mtt.messages, fmt.Sprintf(format, v...))
}

func (mtt *mockTestingT) Messages() []string {
    // Even though this struct is unexported and we could technically use
    // mtt.messages directly it's still a good idea to create a "public"
    // interface so anyone who uses your library in the future knows where
    // to look for usages. It also makes it easier to maintain this in the
    // future because we don't need to guess which "private" fields are
    // being used.
    return mtt.messages
}

// OK! We have our mock *testing.T, time to use it to test our mockable test
// logger.

func TestMockableTestLoggerDebug(t *testing.T) {
    var mockT mockTestingT

    // In this case we want to pass in our mock *testing.T instead of 't' because
    // we're testing the test logger, not just using it.
    mtl := NewMockableTestLogger(&mockT)

    mtl.Debug("message")

    // Make sure our message is in the mock testing.T's messages and is what we expect.
    require.Len(t, mockT.Messages(), 1)
    assert.Equal(t, "[DEBUG(test)] message", mockT.Messages()[0])
}

func TestMockableTestLoggerInfo(t *testing.T) {
    var mockT mockTestingT

    mtl := NewMockableTestLogger(&mockT)

    mtl.Info("message")

    require.Len(t, mockT.Messages(), 1)
    assert.Equal(t, "[INFO(test)] message", mockT.Messages()[0])
}

func TestMockableTestLoggerError(t *testing.T) {
    var mockT mockTestingT

    mtl := NewMockableTestLogger(&mockT)

    mtl.Error("message")

    require.Len(t, mockT.Messages(), 1)
    assert.Equal(t, "[ERROR(test)] message", mockT.Messages()[0])
}

// So now we have our tested, mockable test logger we can use in our tests! Take a look
// at service/service_test.go for that part.