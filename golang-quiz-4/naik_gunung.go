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

	var isNaikSaja bool = true

	var ketinggianSebelumnya int
	var ketinggianSekarang int

	var i int

	scanner.Scan()
	line = scanner.Text()

	angkaString = strings.Split(line, " ")

	ketinggianSebelumnya, _ = strconv.Atoi(angkaString[1])

	for i = 2; i < len(angkaString); i++ {

		ketinggianSekarang, _ = strconv.Atoi(angkaString[i])

		isNaikSaja = isNaikSaja && (ketinggianSekarang > ketinggianSebelumnya)

		ketinggianSebelumnya = ketinggianSekarang
	}

	fmt.Println("test", isNaikSaja)
}
