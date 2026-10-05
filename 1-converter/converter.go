package main

import "fmt"

func main() {
	const (
		usdInEur = 0.92
		usdInRub = 1.60
	)
	eurInRub := usdInRub / usdInEur
	fmt.Println(eurInRub)
}
