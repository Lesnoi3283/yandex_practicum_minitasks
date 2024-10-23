package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func LastModified(dir string, hours int) error {
	// допишите код
	// ...
	return filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if time.Since(info.ModTime()).Hours() > float64(hours) {
			return nil
		} else {
			fmt.Printf("%s %s\n", path, info.ModTime())
			return nil
		}
	})
}

func main() {
	err := LastModified(`/home/vboxuser/Education`, 1)
	if err != nil {
		fmt.Println(err)
	}
}
