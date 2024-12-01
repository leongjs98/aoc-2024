package d2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Puzzle1() {
	reports := getReports("./d2/input")
	exp, err := regexp.Compile("\\d+")
	if err != nil {
		log.Fatal(err)
	}

	numSafeReports := 0
	for _, report := range reports {
		// fmt.Println("report", report)
		levels := []int{}
		levelsStr := exp.FindAllString(report, -1)
		for _, lvlStr := range levelsStr {
			level, err := strconv.Atoi(lvlStr)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, level)
		}
		// fmt.Println("levels", levels)

		isSafe := getReportSafety(levels)
		if isSafe {
			fmt.Println(levels, "is safe")
			numSafeReports++
		}
	}

	fmt.Println("safe:", numSafeReports)
}

func Puzzle2() {
	reports := getReports("./d2/input")
	exp, err := regexp.Compile("\\d+")
	if err != nil {
		log.Fatal(err)
	}

	numSafeReports := 0
	for _, report := range reports {
		levels := []int{}
		levelsStr := exp.FindAllString(report, -1)
		for _, lvlStr := range levelsStr {
			level, err := strconv.Atoi(lvlStr)
			if err != nil {
				log.Fatal(err)
			}
			levels = append(levels, level)
		}

		isSafe := getReportSafety(levels)
		isTolerated := false
		unsafeIdx := 0
		if !isSafe {
			isTolerated = true
			for j := 0; j < len(levels); j++ {
				newLevels := append([]int(nil), levels...)
				newLevels = append(newLevels[:j], newLevels[j+1:]...)
				fmt.Println("test new levels safety", newLevels)
				isSafe = getReportSafety(newLevels)
				fmt.Println()
				if isSafe {
					unsafeIdx = j
					break
				}
			}
		}

		if isSafe && !isTolerated {
			fmt.Println(levels, "is safe without changes")
			numSafeReports++
		} else if isSafe && isTolerated {
			fmt.Println(levels, "is safe after removing index ", unsafeIdx, levels[unsafeIdx])
			numSafeReports++
		} else if !isSafe && isTolerated {
			fmt.Println(levels, "is unsafe regardless of level changes")
		}
	}

	fmt.Println("safe:", numSafeReports)
}

func getReportSafety(levels []int) (isSafe bool) {
	fmt.Println("testing safety of", levels)
	isIncreasing, isDecreasing := false, false
	for i := 0; i < len(levels)-1; i++ {
		if levels[i+1] == levels[i] {
			fmt.Println("unsafe: same level", levels[i], levels[i+1])
			return false
		}
		if i == 0 {
			if levels[i+1] > levels[i] {
				isIncreasing, isDecreasing = true, false
			} else if levels[i+1] < levels[i] {
				isIncreasing, isDecreasing = false, true
			}
		}

		if isIncreasing {
			if levels[i+1] < levels[i] {
				fmt.Println("unsafe: decreased in increasing sequence", levels[i], levels[i+1])
				return false
			}
			if levels[i+1]-levels[i] > 3 || levels[i+1]-levels[i] < 1 {
				fmt.Println("unsafe: increasing difference too large", levels[i], levels[i+1])
				return false
			}
		} else if isDecreasing {
			if levels[i+1] > levels[i] {
				fmt.Println("unsafe: increased in decreasing sequence", levels[i], levels[i+1])
				return false
			}
			if levels[i+1]-levels[i] > -1 || levels[i+1]-levels[i] < -3 {
				fmt.Println("unsafe: decreasing difference too large", levels[i], levels[i+1])
				return false
			}
		}
	}

	return true
}

func getReports(filename string) (lines []string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	lines = []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	return lines
}
