package main

import (
	"os"
	"log"
	"io"
	"bufio"
	"fmt"
	//"strconv"
)

func main() {

	src := "kjevik-temp-celsius-20220318-20230318.csv"
	dest := "kjevik-temp-fahr-20220318-20230318.csv"

	sourceFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(os.Stdin)
        fmt.Println("Type either \"convert\" to make a new file or \"average\" to get the average temperature.")
        scanner.Scan()
        input := scanner.Text()

        if (input == "convert" || input == "Convert") {
                fmt.Println("Test \"convert\" passed.")
        }else if (input == "average" || input == "Average") {
                fmt.Println("Test \"average\" passed.")
        }else {
                fmt.Println("Please type either \"convert\" or \"average\".")
        }
}
