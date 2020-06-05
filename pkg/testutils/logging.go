package testutils

import (
	"testing"

	"github.com/sirupsen/logrus"
)

// NewLog will create a logrus entry with a field that includes the name of the current test.
func NewLog(t *testing.T) *logrus.Entry {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)
	return logrus.NewEntry(logger).WithField("test", t.Name())
}
