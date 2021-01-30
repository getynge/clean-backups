package util_test

import (
	"github.com/getynge/clean-backups/util"
	"testing"
	"time"
)

func TestShouldDeleteFile(t *testing.T) {
	before := time.Date(2021, time.January, 20, 23, 0, 0, 0, time.UTC)
	files := map[string]bool{
		"/home/test/backups/202101192200.tar.gz": true,
		"202101192200.tar.gz":                    true,
		"202001010100.tar.gz":                    true,
		"202101200000.tar.gz":                    false,
		"202101012300.tar.gz":                    false,
		"202101210000.tar.gz":                    false,
	}

	for file, desiredResult := range files {
		actualResult := util.ShouldDeleteFile(file, before)
		if actualResult != desiredResult {
			t.Fatalf("expected delete to be %t for %s", desiredResult, file)
		}
	}
}
