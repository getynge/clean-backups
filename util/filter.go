package util

import (
	"path"
	"time"
)

func ShouldDeleteFile(fqp string, now time.Time) bool {
	file := path.Base(fqp)

	moment, err := time.Parse("200601021504.tar.gz", file)

	if err != nil {
		return false
	}

	difference := now.Sub(moment)

	if difference < 0 || difference < time.Hour*24 {
		return false
	}

	if difference >= time.Hour*24*30 {
		return true
	}

	if moment.Hour() == 23 {
		return false
	}

	return true
}
