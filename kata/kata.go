package kata

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func StockList(listArt []string /*listCat []string*/) string {
	resultSlice := make([]string, 0)

	firstLetter, err := NthItemToSlice(0, listArt)
	if err != nil {
		fmt.Print(err)
	}

	intsSum := make([]int, 0)

	for _, valArt := range listArt {

		var lastInts []string
		toArr := stringToArr(valArt)

		thirdLast, err := NthItemToSlice(len(valArt)-3, toArr)
		if err != nil {
			fmt.Print(err)
		}
		lastInts = append(lastInts, strings.Join(thirdLast, ""))

		secondLast, err := NthItemToSlice(len(valArt)-2, toArr)
		if err != nil {
			fmt.Print(err)
		}
		lastInts = append(lastInts, strings.Join(secondLast, ""))

		last, err := NthItemToSlice(len(valArt)-1, toArr)
		if err != nil {
			fmt.Print(err)
		}
		lastInts = append(lastInts, strings.Join(last, ""))
		total := strings.Join(lastInts, "")
		//fmt.Println("Total is: ", total)
		toI, _ := strconv.Atoi(total)

		intsSum = append(intsSum, toI)
	}
		for index, val := range firstLetter {
		numToS := strconv.Itoa(intsSum[index])
		fmtLetter := fmt.Sprintf("(%s : %s)", string(val), numToS)
		fmt.Print(fmtLetter)
		// First item in slice is Letter
		resultSlice = append(resultSlice, fmtLetter)
	}

	return strings.Join(resultSlice, " - ")
}

func NthItemToSlice(n int, arr []string) ([]string, error) {
	results := make([]string, 0)
	if len(arr) < n {
		err := fmt.Errorf("nthItem: Length of arr is: %v", len(arr))
		log.Fatal(err)
	}

	for _, val := range arr {
		// fmt.Printf("Value: %v\n", val)
		results = append(results, string(val))
	}

	return results, nil //strings.Join(results, ""), nil
}

func stringToArr(s string) []string {

	returnArr := make([]string, 0)

	for _, letter := range s {
		returnArr = append(returnArr, string(letter))
	}

	return returnArr
}