package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	num := len(args)

	if num < 1 {
		fmt.Println("Provide filename as argument.")
		return
	}

	for i := 0; i < num; i++ {
		oldName := args[i]
		fmt.Println(oldName)
		newName := strings.Replace(oldName, " ", "_", -1)
		fmt.Println(newName)

		err := os.Rename(oldName, newName)
		if err != nil {
			log.Fatal(err)
		}
	}

}
