package currtime

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestFtime(t *testing.T) {
	tests := []struct {
		host string
		err  string
	}{
		{
			host: "0.beevik-ntp.pool.ntp.org",
			err:  "",
		},
		{
			host: "",
			err:  "fatal error: address string is empty",
		},
	}

	for i, test := range tests {
		var buf bytes.Buffer
		var errBuf bytes.Buffer

		t.Run(fmt.Sprintf("case: %v", i), func(t *testing.T) {
			ftime(&buf, &errBuf, test.host)

			err := errBuf.String()
			out := buf.String()
			if err != test.err {
				t.Errorf("Got: %v, want: %v", err, test.err)
				return
			}

			if test.err != "" {
				return
			}

			timeLayout := "2006-01-02 15:04:05.999999999 -0700 MST"
			out = strings.Split(out, " m=")[0]

			tout, terr := time.Parse(timeLayout, out)
			if terr != nil {
				t.Errorf("Error parse time\ntime: %v\nerr: %v", out, terr)
				return
			}

			tnow := time.Now()
			if tout.Truncate(time.Second) != tnow.Truncate(time.Second) {
				t.Errorf("Got: %v, want: %v", tnow, tout)
			}
			t.Logf("Got time: %v\n", tout)
			t.Logf("Time now: %v\n", tnow)
		})
	}
}
