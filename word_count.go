package main

import (
	"bufio"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{Use: "wc"}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	rootCmd.AddCommand(countLinesAndWordsCmd)
	countLinesAndWordsCmd.Flags().BoolP("lines", "l", false, "Count lines")
	countLinesAndWordsCmd.Flags().BoolP("words", "w", false, "Count words")

}

var countLinesAndWordsCmd = &cobra.Command{
	Use:   "wc [file]",
	Short: "Count lines or words in a file",
	Args:  cobra.ExactArgs(1),
	Run:   countLinesAndWords,
}

func countLinesAndWords(cmd *cobra.Command, args []string) {
	countLines, _ := cmd.Flags().GetBool("lines")
	countWords, _ := cmd.Flags().GetBool("words")

	fileName := args[0]
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

	if err != nil {
		if errors.Is(err, fs.ErrPermission) {
			fmt.Printf("%s: open: Permission denied \n", fileName)
		} else {
			fmt.Printf("%s: open: Permission denied \n", fileName)
		}
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineCount := 0
	wordCount := 0

	for scanner.Scan() {
		lineStr := scanner.Text()
		if len(strings.TrimSpace(lineStr)) > 0 {
			lineCount++
			if countWords {
				wordCount += countWordsInString(lineStr)
			}
		}
	}

	if scanner.Err() != nil {
		fmt.Printf("%s: read: Unable to read the file \n", fileName)
		return
	}

	if countLines {
		fmt.Printf("    %d %s\n", lineCount, fileName)
	} else {
		fmt.Printf("    %d %s\n", wordCount, fileName)
	}
}

func countWordsInString(str string) int {
	return len(strings.Fields(str))
}
