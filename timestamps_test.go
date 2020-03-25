package main

import (
	"strconv"
	"strings"
	"testing"
	"time"
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
		correctUnix := int64(1451001600000)
		correctUTC := "Fri, 25 Dec 2015 00:00:00"
		ts, err := parseDateString(dateString)

		if err != nil {
			t.Error(err)
		}

		if ts.Unix != correctUnix {
			t.Errorf("Incorrect UNIX timestamp, expecting %d go %d", ts.Unix, correctUnix)
		}

		if !strings.Contains(ts.UTC, correctUTC) {
			t.Errorf("Incorrect UTC timestamp, expecting %q go %q", ts.UTC, correctUTC)
		}
	})

	t.Run("returns current timestamp if empty string is provided", func(t *testing.T) {
		ts, _ := parseDateString("")
		utcNow := time.Now().UTC().Format("Mon, 2 Jan 2006 15:04:05 MST")
		unixNow := time.Now().Unix()

		if !strings.Contains(ts.UTC, utcNow[:20]) {
			t.Error("UTC does not match")
		}

		if !strings.Contains(strconv.FormatInt(ts.Unix, 10), strconv.FormatInt(unixNow, 10)[:5]) {
			t.Error("UNIX does not match")
		}
	})

	t.Run("can parse unix timestamps", func(t *testing.T) {
		dateString := "1450137600000"
		correctUnix := int64(1450137600000)
		correctUTC := "Tue, 15 Dec 2015 00:00:00"
		ts, err := parseDateString(dateString)

		if err != nil {
			t.Error(err)
		}

		if ts.Unix != correctUnix {
			t.Errorf("Incorrect UNIX timestamp, expecting %d got %d", ts.Unix, correctUnix)
		}

		if !strings.Contains(ts.UTC, correctUTC) {
			t.Errorf("Incorrect UTC timestamp, expecting %q got %q", ts.UTC, correctUTC)
		}
	})
}
