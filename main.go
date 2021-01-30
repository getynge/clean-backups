package main

import (
	"fmt"
	"github.com/getynge/clean-backups/util"
	"os"
	"path/filepath"
	"time"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("clean-backups takes exactly one argument, please specify the folder to manage")
		os.Exit(1)
	}

	now := time.Now()
	root := os.Args[1]

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !util.ShouldDeleteFile(path, now) {
			return nil
		}

		return os.Remove(path)
	})

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
