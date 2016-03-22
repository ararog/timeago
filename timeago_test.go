package timeago

import (
	"time"
)

import "testing"

func check(t *testing.T, d time.Duration, result string) {
  start := time.Now()
  end := time.Now().Add(d)
  got := TimeAgo(start, end)
  if got != result {
    t.Errorf("Wrong result: %s", got)
  }
}

func TestHourAgo(t *testing.T) {
  d, error := time.ParseDuration("-1.5h")
  if error == nil {
    check(t, d, "An hour ago")
  }
}

func TestMinuteAgo(t *testing.T) {
  d, error := time.ParseDuration("-1.2m")
  if error == nil {
    check(t, d, "A minute ago")
  }
}
