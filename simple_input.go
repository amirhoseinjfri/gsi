package simple_input

import (
	"bufio"
	"fmt"
	"os"
)

//Func Input getting message as text to show to user
//returnin Scanned Text as string
func Input(message string) string {
	fmt.Printf("%s", message)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}
