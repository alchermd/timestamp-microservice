package main

import (
	"strings"
	"testing"
)

func TestParseDateString(t *testing.T) {
	t.Run("error on invalid dateString", func(t *testing.T) {
		dateString := "hello world"
		_, err := parseDateString(dateString)

		if err == nil {
			t.Errorf("Did not throw an error on %q", dateString)
		}
	})

	t.Run("correct unix and UTC time", func(t *testing.T) {
		dateString := "2015-12-25"
		correctUnix := int64(1451001600)
		correctUTC := "Fri, 25 Dec 2015 00:00:00"
		ts, err := parseDateString(dateString)

		if err != nil {
			t.Error(err)
		}

		if ts.Unix !=  correctUnix{
			t.Errorf("Incorrect UNIX timestamp, expecting %d go %d", ts.Unix, correctUnix)
		}

		if !strings.Contains(ts.UTC, correctUTC) {
			t.Errorf("Incorrect UTC timestamp, expecting %q go %q", ts.UTC, correctUTC)
		}
	})
}