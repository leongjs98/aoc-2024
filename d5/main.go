package d5

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var (
	keyBeforeVal = make(map[int][]int)
	keyAfterVal  = make(map[int][]int)
)

func Puzzle1() {
	rules := getLines("./d5/rules")
	// fmt.Println("Rules:")
	// for _, rule := range rules {
	// 	fmt.Println(rule)
	// }

	keyBeforeVal, keyAfterVal = genRuleMaps(rules)
	// fmt.Println(keyBeforeVal)
	// fmt.Println(keyAfterVal)

	updatesStr := getLines("./d5/updates")
	updates := getUpdates(updatesStr)
	sums := 0
	for _, nums := range updates {
		isCorrect := true
		for i := range nums {
			numsAfter := keyBeforeVal[nums[i]]
			numsBefore := keyAfterVal[nums[i]]
			// fmt.Println(numsBefore, numsAfter)
			// compare with numbers before
			for j := 0; j < i; j++ {
				for _, numAfter := range numsAfter {
					if nums[j] == numAfter {
						isCorrect = false
						break
					}
				}
				if !isCorrect {
					break
				}
			}
			// compare with numbers after
			for j := i + 1; j < len(nums)-1; j++ {
				for _, numBefore := range numsBefore {
					if nums[j] == numBefore {
						isCorrect = false
						break
					}
				}
				if !isCorrect {
					break
				}
			}
			if !isCorrect {
				break
			}
		}
		if isCorrect {
			fmt.Println("correct!:", nums, nums[(len(nums)/2)])
			sums += nums[(len(nums) / 2)]
		}
	}

	fmt.Println(sums)
}

func Puzzle2() {
	rules := getLines("./d5/rules")

	keyBeforeVal, keyAfterVal = genRuleMaps(rules)
	fmt.Println("keyBeforeVal", keyBeforeVal)
	fmt.Println("keyAfterVal", keyAfterVal)

	updatesStr := getLines("./d5/updates")
	updates := getUpdates(updatesStr)

	incorrectUpdates := [][]int{}
	for _, nums := range updates {
		isCorrect := true
		for i := range nums {
			numsAfter := keyBeforeVal[nums[i]]
			numsBefore := keyAfterVal[nums[i]]
			// fmt.Println(numsBefore, numsAfter)
			// compare with numbers before
			for j := 0; j < i; j++ {
				for _, numAfter := range numsAfter {
					if nums[j] == numAfter {
						isCorrect = false
						break
					}
				}
				if !isCorrect {
					break
				}
			}
			// compare with numbers after
			for j := i + 1; j < len(nums)-1; j++ {
				for _, numBefore := range numsBefore {
					if nums[j] == numBefore {
						isCorrect = false
						break
					}
				}
				if !isCorrect {
					break
				}
			}
			if !isCorrect {
				break
			}
		}
		if !isCorrect {
			incorrectUpdates = append(incorrectUpdates, nums)
		}
	}

	sums := 0
	for _, update := range incorrectUpdates {
		fmt.Println("original:", update)
		update = bubblesort(update)
		fmt.Println("corrected:", update, update[(len(update)/2)])
		sums += update[(len(update) / 2)]
	}
	fmt.Println(sums, len(incorrectUpdates))

}

func bubblesort(update []int) []int {
	for j := len(update) - 1; j > 0; j-- {
		for i := 0; i < j; i++ {
			if xIsAfterY(update[i], update[i+1]) {
				tmp := update[i]
				update[i] = update[i+1]
				update[i+1] = tmp
			}
		}
	}

	return update
}

func xIsAfterY(x int, y int) bool {
	if _, ok := keyAfterVal[x]; !ok {
		return false
	}
	for _, num := range keyAfterVal[x] {
		if num == y {
			return true
		}
	}
	return false
}

func getUpdates(updatesStr []string) [][]int {
	updates := [][]int{}
	for _, update := range updatesStr {
		numExp := regexp.MustCompile("\\d+")
		numStrs := numExp.FindAllString(update, -1)
		nums := []int{}
		for _, str := range numStrs {
			num, err := strconv.Atoi(str)
			if err != nil {
				panic(err)
			}
			nums = append(nums, num)
		}
		updates = append(updates, nums)
	}
	return updates
}

func genRuleMaps(rules []string) (keyBeforeVal, keyAfterVal map[int][]int) {
	keyBeforeVal = make(map[int][]int)
	keyAfterVal = make(map[int][]int)
	for _, rule := range rules {
		numExp := regexp.MustCompile("\\d+")
		numStrs := numExp.FindAllString(rule, -1)
		num1, err := strconv.Atoi(numStrs[0])
		if err != nil {
			panic(err)
		}
		num2, err := strconv.Atoi(numStrs[1])
		if err != nil {
			panic(err)
		}
		keyBeforeVal[num1] = append(keyBeforeVal[num1], num2)
		keyAfterVal[num2] = append(keyAfterVal[num2], num1)
	}

	return keyBeforeVal, keyAfterVal
}

func getLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	return lines
}
