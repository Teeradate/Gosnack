package main

import (
	"fmt"
)

func main() {

	for i:=1;i<101;i++ {
		fmt.Printf("Decimal : %d, Roman : %s\n",i,intToRoman(i))
	}
}

func intToRoman(num int) string {
	val := []int{
			100, 90, 50, 40, 10, 9, 5, 4, 1,
			}

	sym := []string{
		    "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I",
			}
	
	result := ""
	i := 0
	for num > 0 {
		j := num/val[i]
		for k:=0 ; k < j; k++ {
			result += sym[i]
			num -= val[i]
		}
		i++
	}
	return result
}