package main

import (
	"fmt"
	"log"
	"os"
)

func fileTraversal(path string) {
	fi, err := os.ReadDir(path)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	for _, f := range fi {
		fmt.Println(f.Name())
		if f.IsDir() {
			fmt.Print(" - Directory\n")
			fileTraversal(path + "/" + f.Name())
		}
	}
}
