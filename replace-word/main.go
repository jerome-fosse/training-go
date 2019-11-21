package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	newtext, occ, lines, err := findAndReplace("wikigo.txt", "go", "Python")
	if err != nil {
		fmt.Printf("An error occured : %s\n", err)
		return
	}

	showSummary("go", occ, lines)
	fmt.Println(newtext)
}

func findAndReplace(file string, oldstr string, newstr string) (newText string, occ int, lines []int, err error) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return
	}

	i := 1
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		if nb, line := processLine(scanner.Text(), oldstr, newstr); nb > 0 {
			occ += nb
			lines = append(lines, i)
			newText += line + "\n"
		} else {
			newText += scanner.Text() + "\n"
		}

		i++
	}

	return
}

func processLine(line string, oldstr string, newstr string) (occ int, newline string) {
	if nb, s := replaceAllWithSuffix(line, " ", strings.ToLower(oldstr), strings.ToLower(newstr)); nb > 0 {
		occ += nb
		newline = s
	}
	if nb, s := replaceAllWithSuffix(line, ".", strings.ToLower(oldstr), strings.ToLower(newstr)); nb > 0 {
		occ += nb
		newline = s
	}
	if nb, s := replaceAllWithSuffix(line, ",", strings.ToLower(oldstr), strings.ToLower(newstr)); nb > 0 {
		occ += nb
		newline = s
	}
	if nb, s := replaceAllWithSuffix(line, " ", firstLetterUpper(oldstr), firstLetterUpper(newstr)); nb > 0 {
		occ += nb
		newline = s
	}
	if nb, s := replaceAllWithSuffix(line, ".", firstLetterUpper(oldstr), firstLetterUpper(newstr)); nb > 0 {
		occ += nb
		newline = s
	}
	if nb, s := replaceAllWithSuffix(line, ",", firstLetterUpper(oldstr), firstLetterUpper(newstr)); nb > 0 {
		occ += nb
		newline = s
	}

	return
}

func replaceAllWithSuffix(line, suffix, oldstr, newstr string) (occ int, newline string) {
	if c := strings.Count(line, oldstr+suffix); c > 0 {
		occ += c
		newline = strings.ReplaceAll(line, oldstr+suffix, newstr+suffix)
	}

	return
}

func firstLetterUpper(s string) string {
	return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
}

func showSummary(str string, occ int, lines []int) {
	fmt.Println("=== Summary ===")
	fmt.Printf("Number of occurences of %s : %d\n", str, occ)
	fmt.Printf("Number of lines : %d\n", len(lines))
	fmt.Printf("Lines : %v\n", lines)
	fmt.Println("=== End of Summary ===")
	fmt.Println("")
}
