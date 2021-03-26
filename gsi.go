package gsi

import (
	"bufio"
	"fmt"
	"os"
)

//function Input Returning user input as string
func Input(a ...interface{}) string {
	fmt.Printf("%v", a...)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
