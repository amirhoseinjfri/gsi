package gsi

import (
	"bufio"
	"fmt"
	"os"
)

//function Input Returning user input as string
//you will here beep in cmd or powershell
func Input(s bool, a ...interface{}) string {
	if s {
		fmt.Printf("\a")
	}
	fmt.Printf("%v", a...)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
