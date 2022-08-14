package main

import (
	"flag"
	"fmt"
	"strconv"
)

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	i, _ := strconv.Atoi(arg)

	fmt.Println("¥", taxIncluded(i))
}

func taxIncluded(i int) float64 {
	f := float64(i)
	consumptionTax := f * 0.1
	taxIncludedPrice := f + consumptionTax

	return taxIncludedPrice
}
