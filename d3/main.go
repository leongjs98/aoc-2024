package d3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Puzzle1() {
	lines := getLines("./d3/input")

	realInstExp, err := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)")
	if err != nil {
		log.Fatal(err)
	}

	insts := []string{}
	for _, line := range lines {
		tmpInsts := realInstExp.FindAllString(line, -1)
		insts = append(insts, tmpInsts...)
	}

	numExp, err := regexp.Compile("\\d+")
	sum := 0
	for _, inst := range insts {
		numsStr := numExp.FindAllString(inst, -1)
		num1, _ := strconv.Atoi(numsStr[0])
		num2, _ := strconv.Atoi(numsStr[1])
		sum += num1 * num2
	}

	fmt.Println(sum)
}

func Puzzle2() {
	lines := getLines("./d3/input")
  fmt.Println(lines)

	instExp, err := regexp.Compile("mul\\(\\d{1,3},\\d{1,3}\\)|don't\\(\\)|do\\(\\)")
	if err != nil {
		log.Fatal(err)
	}

	insts := []string{}
	for _, line := range lines {
		tmpInsts := instExp.FindAllString(line, -1)
		insts = append(insts, tmpInsts...)
	}

	isDo := true
	newInsts := []string{}
	for _, inst := range insts {
		if inst == "don't()" {
			isDo = false
		} else if inst == "do()" {
			isDo = true
		} else if isDo {
			newInsts = append(newInsts, inst)
		}
	}

	fmt.Println(newInsts)
	numExp, err := regexp.Compile("\\d+")
	sum := 0
	for _, inst := range newInsts {
		numsStr := numExp.FindAllString(inst, -1)
		num1, _ := strconv.Atoi(numsStr[0])
		num2, _ := strconv.Atoi(numsStr[1])
		sum += num1 * num2
	}

	fmt.Println(sum)
}

func getLines(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
