// Package currtime  prints the exact current time.
package currtime

import (
	"fmt"
	"io"
	"os"

	"github.com/beevik/ntp"
)

// Ftime writes the current time time.Time to w.
func Ftime(w io.Writer, host string) {
	err := ftime(w, os.Stderr, host)
	if err != nil {
		os.Exit(1)
	}
}

func ftime(w io.Writer, errw io.Writer, host string) error {
	time, err := ntp.Time(host)
	if err != nil {
		fmt.Fprint(errw, "fatal error: ", err)
		return err
	}

	fmt.Fprint(w, time.String())
	return nil
}
