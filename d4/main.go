package d4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func Puzzle1() {
	lines := []string{}

	chars := getChars("./d4/input")
	fmt.Println("original")
	for _, line := range chars {
		lines = append(lines, strings.Join(line, ""))
	}

	fmt.Println("\nvertical transposed:")
	vertTpChars := getVertTranspose(chars)
	for _, line := range vertTpChars {
		lines = append(lines, strings.Join(line, ""))
	}

	fmt.Println("\forwardslash transposed:")
	forwSlashTpChars := getForwslahTranspose(chars)
	for _, line := range forwSlashTpChars {
		lines = append(lines, strings.Join(line, ""))
	}

	fmt.Println("\nbackslash transposed:")
	backslashTpChars := getBackslahTranspose(chars)
	for _, line := range backslashTpChars {
		lines = append(lines, strings.Join(line, ""))
	}

	count := countWordSearch(lines)
	fmt.Println(count)
}

func Puzzle2() {
	chars := getChars("./d4/input")

	count := 0

	for i := range chars {
		for j := range chars[i] {
			if chars[i][j] == "A" {
				if !(i-1 < 0 || i+1 > len(chars)-1 || j+1 > len(chars[i])-1 || j-1 < 0) {
					topLeft := chars[i-1][j-1]
					topRight := chars[i-1][j+1]
					bottomLeft := chars[i+1][j-1]
					bottomRight := chars[i+1][j+1]
					isforwSlashMAS := (topLeft == "M" && bottomRight == "S") || (topLeft == "S" && bottomRight == "M")
					isbackSlashMAS := (bottomLeft == "M" && topRight == "S") || (bottomLeft == "S" && topRight == "M")
					if isforwSlashMAS && isbackSlashMAS {
            count++
					}
				}
			}
		}
	}

	fmt.Println(count)
}

func countWordSearch(lines []string) (count int) {
	expXMAS, err := regexp.Compile("XMAS")
	if err != nil {
		panic(err)
	}
	expSAMX, err := regexp.Compile("SAMX")
	if err != nil {
		panic(err)
	}
	count = 0
	for _, line := range lines {
		words := expXMAS.FindAllString(line, -1)
		count += len(words)
		words = expSAMX.FindAllString(line, -1)
		count += len(words)
	}
	return count
}

func getVertTranspose(chars [][]string) (newChars [][]string) {
	for i := range chars {
		tmpChars := []string{}
		for j := range chars[i] {
			tmpChars = append(tmpChars, chars[j][i])
		}
		newChars = append(newChars, tmpChars)
	}
	return newChars
}

func getForwslahTranspose(chars [][]string) (newChars [][]string) {
	// top left half
	for i := range chars[0] {
		tmpChars := []string{chars[0][i]}
		for j := 1; j < len(chars) && i-j >= 0; j++ {
			tmpChars = append(tmpChars, chars[j][i-j])
		}
		newChars = append(newChars, tmpChars)
	}

	// bottom right half
	for i := 1; i < len(chars); i++ {
		tmpChars := []string{chars[i][len(chars)-1]}
		for j := 1; i+j < len(chars) && len(chars)-1-j >= 0; j++ {
			tmpChars = append(tmpChars, chars[i+j][len(chars)-1-j])
		}
		newChars = append(newChars, tmpChars)
	}

	return newChars
}

func getBackslahTranspose(chars [][]string) (newChars [][]string) {
	// top right half
	for i := range chars[0] {
		tmpChars := []string{chars[0][i]}
		for j := 1; j < len(chars) && i+j < len(chars[j]); j++ {
			tmpChars = append(tmpChars, chars[j][i+j])
		}
		newChars = append(newChars, tmpChars)
	}

	// bottom left half
	for i := 1; i < len(chars); i++ {
		tmpChars := []string{chars[i][0]}
		for j := 1; i+j < len(chars) && j < len(chars); j++ {
			tmpChars = append(tmpChars, chars[i+j][j])
		}
		newChars = append(newChars, tmpChars)
	}

	return newChars
}

func getChars(filename string) (chars [][]string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		chars = append(chars, strings.Split(fileScanner.Text(), ""))
	}
	return chars
}
