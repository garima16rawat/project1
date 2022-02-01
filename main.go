package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating for pizza ")
	input, _ := reader.ReadString('\n') //comma ok or comma error syntex
	fmt.Println("thanks for rating", input)
	fmt.Printf("type of rating is: %T ", input)
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("add one to rating:", numrating+1)
	}

}
													