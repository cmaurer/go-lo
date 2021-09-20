package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

//go:generate go run

func main() {
	// Get all Subdirs
	dirEntries, err := os.ReadDir("codegen/")
	if err != nil {
		panic(err)
	}
	for _, dirEntry := range dirEntries {
		if dirEntry.Type().IsDir() {
			if strings.LastIndex(dirEntry.Name(), "types") == -1 {
				path := fmt.Sprintf("./%s/%s", "codegen", dirEntry.Name())
				fmt.Println(path)
				cmd := exec.Command("go", "run", path, path)
				var out bytes.Buffer
				cmd.Stdout = &out

				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}
