package time

import (
  "time"
  "strings"
)

func Time() time.Time{
  now := time.Now()
  now2 := strings.Replace(now.Format(time.RFC3339), "T", " ", 1)
  now3 := strings.Replace(now2, "+09:00","", 1)
  now4 := strings.Replace(now3, "Z", "", 1)
  layout := "2006-01-02 15:04:05"
  t, _ := time.Parse(layout, now4)
  t2 := t.Add(9 * time.Hour)
	return t2
}
