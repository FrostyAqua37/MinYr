package main

import (
	"os"
	"log"
	"io/ioutil"
	"bufio"
	"fmt"
	"strings"
	//"strconv"
)

func main() {

	src := "kjevik-temp-celsius-20220318-20230318.csv"
	dest := "kjevik-temp-fahr-20220318-20230318.csv"

	//åpner filen. 
	sourceFile, err := os.Open(src)
	if err != nil {
		log.Fatal(err)
	}
	defer sourceFile.Close()

	//Leser innholdet til filen, en linje om gangen og deler linjene mellom 4 ulike deler. 
	scanner := bufio.NewScanner(sourceFile)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ";")

		fmt.Println(items[3])
	}

	//Lager filen der innholdet kan kopieres til. 
	destinationFile, err := os.Create(dest)
	if err != nil {
		log.Fatal(err)
	}
	defer destinationFile.Close()

	content, err = ioutil.WriteFile(destinationFile, scanner.Text(), 0644)
	if err != nil {
		log.Fatal(err)
	}

	//Tar input fra bruker og gjør om kommandoer basert på input. 
	inputScanner := bufio.NewScanner(os.Stdin)
        fmt.Println("Type either \"convert\" to make a new file or \"average\" to get the average temperature.")
        inputScanner.Scan()
        input := inputScanner.Text()

        if (input == "convert" || input == "Convert") {
                fmt.Println("Test \"convert\" passed.")
        }else if (input == "average" || input == "Average") {
                fmt.Println("Test \"average\" passed.")
        }else {
                fmt.Println("Please type either \"convert\" or \"average\".")
        }
}
