package main

import (
	"bytes"
	"flag"
	"io"
	"log/syslog"
	"os"
	"os/exec"

	"github.com/kr/pty"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	logfile string
	rsyslog string
)

func main() {
	flag.StringVar(&logfile, "logfile", "./log.txt", "Path to the log file.")
	flag.StringVar(&rsyslog, "syslog", "127.0.0.1:512", "Remote syslog address.")

	flag.Parse()

	cmd := exec.Command("/bin/bash")
	tty, err := pty.Start(cmd)
	if err != nil {
		panic(err)
	}
	defer tty.Close()

	oldState, err := terminal.MakeRaw(0)
	if err != nil {
		panic(err)
	}
	defer terminal.Restore(0, oldState)

	file, err := os.OpenFile(logfile, os.O_RDWR|os.O_APPEND, 0775)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	logwriter, err := syslog.Dial("udp", rsyslog, syslog.LOG_DEBUG, "rootsh")
	if err != nil {
		panic(err)
	}
	defer logwriter.Close()

	l := &Logger{logwriter: logwriter, buffer: &bytes.Buffer{}}

	mw := io.MultiWriter(os.Stdout, file, l)

	go io.Copy(tty, os.Stdin)
	io.Copy(mw, tty)
}
