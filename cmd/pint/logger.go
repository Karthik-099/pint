package main

import (
	"fmt"

	"github.com/cloudflare/pint/internal/log"
)

func initLogger(level, format string, noColor bool) error {
	l, err := log.ParseLevel(level)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid log level", level)
	}

	f, err := log.ParseFormat(format)
	if err != nil {
		return fmt.Errorf("'%s' is not a valid log format", format)
	}

	log.Setup(l, f, noColor)

	return nil
}
