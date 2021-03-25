package gsi

import (
	"bufio"
	"fmt"
	"os"
)

//function Input Returning user input as string
func Input(message string) string {
	fmt.Printf("%s", message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
