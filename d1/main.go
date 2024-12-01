package d1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Puzzle1() {
	leftList, rightList := getLists("./d1/input")

	ascLeftList := ascendingBubbleSort(leftList)
	ascRightList := ascendingBubbleSort(rightList)

	distList := []int{}
	if len(ascLeftList) == len(ascRightList) {
		for i := range ascLeftList {
			dist := ascLeftList[i] - ascRightList[i]
			if dist < 0 {
				dist = -1 * dist
			}
			distList = append(distList, dist)
		}
	} else {
		log.Fatalf("lists not equal len")
	}

	sum := 0
	for _, dist := range distList {
		sum += dist
	}
	fmt.Println(sum)
}

func Puzzle2() {
	leftList, rightList := getLists("./d1/input")

	ascLeftList := ascendingBubbleSort(leftList)
	ascRightList := ascendingBubbleSort(rightList)

	leftNumAppearanceOnRight := make(map[int]int)
	if len(ascLeftList) == len(ascRightList) {
		for i := range ascLeftList {
			if _, ok := leftNumAppearanceOnRight[ascLeftList[i]]; !ok {
				leftNumAppearanceOnRight[ascLeftList[i]] = 0
				for _, rightNum := range ascRightList {
					if rightNum == ascLeftList[i] {
						leftNumAppearanceOnRight[ascLeftList[i]]++
					}
				}
			}
		}
	} else {
		log.Fatalf("lists not equal len")
	}

	similarityScores := 0
	for _, leftNum := range ascLeftList {
		similarityScores += leftNum * leftNumAppearanceOnRight[leftNum]
	}
	fmt.Println(similarityScores)
}

func getLists(filename string) (leftList, rightList []int) {
	readFile, err := os.Open(filename)
	defer readFile.Close()
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	leftExp, err := regexp.Compile("^\\d+")
	if err != nil {
		log.Fatal(err)
	}
	rightExp, err := regexp.Compile("\\d+$")
	if err != nil {
		log.Fatal(err)
	}

	for fileScanner.Scan() {
		leftList = append(leftList, getNum(leftExp, fileScanner.Text()))
		rightList = append(rightList, getNum(rightExp, fileScanner.Text()))
	}

	return leftList, rightList
}

func getNum(regExp *regexp.Regexp, text string) int {
	str := regExp.FindString(text)
	num, err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func ascendingBubbleSort(list []int) []int {
	sortedList := make([]int, len(list))
	copy(sortedList, list)

	for i := 1; len(sortedList)-i > 1; i++ {
		for j := 0; j < len(sortedList)-i; j++ {
			if sortedList[j] > sortedList[j+1] {
				tmpNum := sortedList[j+1]
				sortedList[j+1] = sortedList[j]
				sortedList[j] = tmpNum
			}
		}
	}

	return sortedList
}
