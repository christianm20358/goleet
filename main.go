package main

import (
	"fmt"
	"strings"
)

var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
var c = []string{"A", "B", "C", "D"}

func StockList(listArt []string, listCat []string) string {
	// your code
	//var result string

	slice := make([]string, 0)

	m := make(map[int]string)

	for key, val := range listArt {
		slice = append(slice, val[len(val)-1:])
		fmt.Println(slice)
		stringFromSlice := strings.Join(slice, "")
		m[key] = stringFromSlice
		return stringFromSlice
	}
	return ""
}

func main() {
	StockList(b, c)

	fmt.Println()
}
