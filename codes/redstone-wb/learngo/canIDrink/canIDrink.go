package canIDrink

import "fmt"

// func CanIDrink(age int) bool {
// 	if age < 18 {
// 		answer := false
// 		fmt.Println(answer)
// 		return false
// 	}
// 	answer := true
// 	fmt.Println(answer)
// 	return true

// }

// Variable Expiration
// func CanIDrink(age int) bool {
// 	if koreanAge := age + 2; koreanAge < 18 {
// 		answer := false
// 		fmt.Println(answer)
// 		return false
// 	}
// 	answer := true
// 	fmt.Println(answer)
// 	return true

// }

//Switch
func CanIDrink(age int) bool {
	switch {
	case age < 18:
		fmt.Println("false!")
	case age > 18:
		fmt.Println("true!")
	}
	return true
}
