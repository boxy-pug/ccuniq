package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	var file *os.File
	var err error
	var countCol bool
	flag.BoolVar(&countCol, "c", false, "enable count column")

	flag.Parse()

	args := flag.Args()

	if len(args) == 0 || args[0] == string('-') {
		file = os.Stdin
	} else {
		file, err = os.Open(string(args[0]))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	removeDuplicateAdjecent(file)

}

func removeDuplicateAdjecent(file *os.File) {
	scanner := bufio.NewScanner(file)
	var lastLine string

	for scanner.Scan() {
		if scanner.Text() == lastLine {
			continue
		} else {
			lastLine = scanner.Text()
			fmt.Println(scanner.Text())
		}

	}

}
