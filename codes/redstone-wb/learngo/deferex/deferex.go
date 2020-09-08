// package main
package deferex

import (
	"fmt"
	"strings"
)

func LenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done!")
	length = len(name)
	uppercase = strings.ToUpper(name)
	return
}
