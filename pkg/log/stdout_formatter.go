package log

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/Sirupsen/logrus"
)

// StdOutFormatter struct
type StdOutFormatter struct {
}

// Format change the default output format to incluse the log level
func (f *StdOutFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	b := &bytes.Buffer{}
	fmt.Fprintf(b, "[%s] - %s\n", strings.ToUpper(entry.Level.String()), entry.Message)
	return b.Bytes(), nil
}
