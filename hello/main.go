package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	myString := "Hello from go."
	fmt.Println(myString)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please enter text:")
	input, _ := reader.ReadString('\n')
	fmt.Println("you entered:", input)

	fmt.Print("Please enter number:")
	numInput, _ := reader.ReadString('\n')
	aFloat, error := strconv.ParseFloat(strings.TrimSpace(numInput), 64)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(aFloat)
	}
}
