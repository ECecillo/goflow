package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("goflow>")
		scanner.Scan()
		err := scanner.Err()
		if err != nil {
			log.Fatal(err)
		}

		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		if line == "exit" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		fmt.Printf("So you said : %s \n", line)
	}
}
