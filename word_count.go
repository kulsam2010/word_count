package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"os"
	"strings"
)

func main() {

	countLines := flag.Bool("l", false, "line count")
	countWords := flag.Bool("w", false, "word count")

	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Println("Usage: word_count [-w | -l] <filename>")
		return
	}

	fileName := flag.Arg(0)
	fileStat, err := os.Stat(fileName)

	if err != nil {
		if errors.Is(err, os.ErrExist) {
			fmt.Printf("%s: open: No such file or directory", fileName)
			return
		}

		if fileStat.IsDir() {
			fmt.Printf("%s: : read: Is a directory", fileName)
			return
		}
	}

	file, err := os.Open(fileName)

	if errors.Is(err, fs.ErrPermission) {
		fmt.Printf("%s: open: Permission denied \n", fileName)
	} else {
		fmt.Printf("%s: open: Permission denied \n", fileName)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	wordCount := 0

	for scanner.Scan() {
		lineCount++

		if *countWords {
			words := strings.Fields(scanner.Text())
			wordCount += len(words)
		}
	}

	if scanner.Err() != nil {
		fmt.Printf("%s: read: Unable to read the file \n", fileName)
		return
	}

	if *countLines {
		fmt.Printf("%d %s\n", lineCount, fileName)
	} else {
		fmt.Printf("%d %s\n", wordCount, fileName)
	}

}
