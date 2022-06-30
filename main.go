package main

import (
	"fmt"

	"github.com/christianm20358/goleet/kata"
)

var b = []string{"BBAR 150", "CDXE 515", "BKWR 250", "BTSQ 890", "DRTY 600"}
var c = []string{"A", "B", "C", "D"}

func main() {

	returnVal := kata.StockList(b)
	fmt.Println(returnVal)

}
