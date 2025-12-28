package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var line string
	var angkaString []string
	var counter int = 0
	var bilangan int
	var i int

	scanner.Scan()
	line = scanner.Text()

	angkaString = strings.Split(line, " ")

	for i = 1; i < len(angkaString); i++ {

		bilangan, _ = strconv.Atoi(angkaString[i])

		counter = counter + (bilangan%10)/9
	}
	fmt.Println(counter)
}
