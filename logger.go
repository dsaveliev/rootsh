package main

import (
	"bytes"
	"log/syslog"
)

type Logger struct {
	logwriter *syslog.Writer
	buffer    *bytes.Buffer
}

func (l *Logger) Write(p []byte) (n int, err error) {
	n = len(p)
	l.buffer.Write(p)

	if bytes.Contains(p, []byte{'\n'}) {
		for _, line := range bytes.Split(l.buffer.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			_, _ = l.logwriter.Write(line)
		}
		l.buffer.Reset()
	}

	return n, nil
}
