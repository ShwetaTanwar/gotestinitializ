package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	welcome := "welcome to user input"
	Println(welcome)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Ente the rating for our pizza")

	input, _ := reader.ReadString


}
