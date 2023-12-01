package cmd

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"

	"github.com/spf13/cobra"
)

// code1 represents the first sub-command
var code1 = &cobra.Command{
	Use:   "code1",
	Short: "code",
	Long:  `This is the code1 sub-command.`,
	Run:   code1Func,
}

func toNumber(s string) int {
	if s == "" {
		return 0
	}
	start := 0
	end := len(s) - 1
	for start <= end {
		startChar := s[start]
		endChar := s[end]

		if !unicode.IsDigit(rune(startChar)) {
			start++
			continue
		}

		if !unicode.IsDigit(rune(endChar)) {
			end--
			continue
		}

		if unicode.IsDigit(rune(startChar)) && unicode.IsDigit(rune(endChar)) {
			break
		}
	}
	if start > end {
		return 0
	}
	result := ""
	if unicode.IsDigit(rune(s[start])) {
		result += string(s[start])
	}
	if unicode.IsDigit(rune(s[end])) {
		result += string(s[end])
	}
	if len(result) == 0 {
		return 0
	}

	number, err := strconv.Atoi(result)
	if err != nil {
		// Handle error
		fmt.Println(err)
		return 0
	}

	return number

}

func code1Func(cmd *cobra.Command, args []string) {
	filename := "./cmd/code1.txt"
	lines := readLines(filename)
	total := 0
	for _, v := range lines {
		n := toNumber(v)
		d := strconv.Itoa(n) + " <= " + v
		fmt.Println(d)
		total += n
	}
	fmt.Println(total)
}

func readLines(filename string) []string {
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Print(err)
	}
	str := string(b)
	lines := strings.Split(str, "\n")
	return lines
}

func init() {
	// Add sub-commands to the root command
	rootCmd.AddCommand(code1)
}
