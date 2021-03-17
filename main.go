package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	args := os.Args[1:]
	num := len(args)

	if num < 1 {
		files, err := filepath.Glob("*")
		if err != nil {
			log.Fatal(err)
		}
		args = files
		num = len(args)
	}

	for i := 0; i < num; i++ {
		oldName := args[i]
		newName := strings.Replace(oldName, " ", "_", -1)
		if newName == oldName {
			continue
		}
		fmt.Println(oldName + " > " + newName)

		err := os.Rename(oldName, newName)
		if err != nil {
			log.Fatal(err)
		}
	}

}
