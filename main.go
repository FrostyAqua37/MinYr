package main

import (
        "os"
        "log"
//      "io/ioutil"
        "bufio"
        "fmt"
        "strings"
        "strconv"

	"github.com/FrostyAqua37/funtemps/conv"
)

func main() {

        src := "kjevik-temp-celsius-20220318-20230318.csv" //Sourcfile is assigned to "src".
//        dest := "kjevik-temp-fahr-20220318-20230318.csv" //Destinationfile is assigned to "dest".


        sourceFile, err := os.Open(src) //Opens the sourcfile.
        if err != nil { //Checks if there is an error or not. 
                log.Fatal(err)
        }
        defer sourceFile.Close() //Command to close the sourcfile last. 

        //Leser innholdet til filen, en linje om gangen og deler linjene mellom 4 ulike deler.
       	/*scanner := bufio.NewScanner(sourceFile)
        for scanner.Scan() {
                line := scanner.Text()
                items := strings.Split(line, ";")

                fmt.Println(items[3])
        }*/

        //Lager filen der innholdet kan kopieres til.
  /*      destinationFile, err := os.Create(dest)
        if err != nil {
                log.Fatal(err)
        }
        defer destinationFile.Close() */

        //Makes a new buffered scanner of input called "inputScanner".
        inputScanner := bufio.NewScanner(os.Stdin)
        fmt.Println("Type either \"convert\" to make a new file or \"average\" to get the average temperature.")
        inputScanner.Scan()
        input := inputScanner.Text() //Takes the input and puts it in "input".

        if (input == "convert" || input == "Convert") { //Checks if the input is either "convert" or "Convert"
		inputScanner02 := bufio.NewScanner(os.Stdin) //New buffered scanner called "inputScanner02".
                fmt.Println("Type either \"Celsius\" or \"Fahrenheit\" to chose output temperature.")
			inputScanner02.Scan()
			input02 := inputScanner02.Text() //Takes input from "inputScanner02" and puts it in "input02".

				if (input02 == "Celsius" || input02 == "celsius" || input02 == "c" || input02 == "C") { //Checks if the input is one of the options listed.
					scanner := bufio.NewScanner(sourceFile) //New scanner to "read" the sourcefile. 
					for scanner.Scan() { //Loops through the file line by line and prints it out. .
						line := scanner.Text()
						items := strings.Split(line, ";") //Divides each line when a ";" appears in the line. 

						fmt.Println(items)
					}

				}else if (input02 == "Fahrenheit" || input02 == "fahrenheit" || input02 == "f" || input02 == "F") { //Checks if the input is one of the options listed. 
					scanner02 := bufio.NewScanner(sourceFile)
					scanner02.Scan()
					var fahrenheit float64

					for scanner02.Scan() {
						line02 := scanner02.Text()
						items02 := strings.Split(line02, ";")

						f, err := strconv.ParseFloat(items02[3], 64)
						if err != nil {
							log.Fatal(err)
						}
						fahrenheit = conv.CelsiusToFahrenheit(f)

						fmt.Println(items02[0], items02[1], items02[2], fahrenheit)
					}
				}else {
					fmt.Println("Please type either \"Celsius\" or \"Fahrenheit\"")
				}

        }else if (input == "average" || input == "Average") {
               inputScanner := bufio.NewScanner(os.Stdin)
	       fmt.Println("Type either \"Celsius\" or \"Fahrenheit\" to chose the average temperature")
			inputScanner.Scan()
			input := inputScanner.Text()

			if (input == "Celsius" || input == "celsius" || input =="c" || input == "C") {
				fileScanner := bufio.NewScanner(sourceFile)
				fileScanner.Scan()
				lineCount := 1

				for fileScanner.Scan() {
//					line := fileScanner.Text()
	//				items := strings.Split(line, ";")

				/*	c, err := strconv.ParseInt(items[3], 10, 64)
					if err != nil {
						log.Fatal(err)
					} */
					lineCount++

				}
				fmt.Println("Number of lines in file:", lineCount)



			}else if (input == "Fahrenheit" || input == "fahrenheit" || input == "f" || input == "F") {
				fileScanner := bufio.NewScanner(sourceFile)
				fileScanner.Scan()
				lineCount := 1

				for fileScanner.Scan() {
					lineCount++
				}
				fmt.Println("Number of lines in file:", lineCount)

			}else {
				fmt.Println("Please type either \"Celsius \" or \"Fahrenheit\".")
			}

        }else {
                fmt.Println("Please type either \"convert\" or \"average\".")
        }
}

