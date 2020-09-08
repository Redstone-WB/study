package superAdd

import "fmt"

// 순차적으로 Print
// func SuperAdd(numbers ...int) int {
// 	for index, number := range numbers {
// 		fmt.Println(index, number)
// 	}
// 	return 1
// }

func SuperAdd(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	fmt.Println(total)
	return 1
}
