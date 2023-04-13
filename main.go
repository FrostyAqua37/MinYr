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

        src := "kjevik-temp-celsius-20220318-20230318.csv"
        dest := "kjevik-temp-fahr-20220318-20230318.csv"

        //  pner filen.
        sourceFile, err := os.Open(src)
        if err != nil {
                log.Fatal(err)
        }
        defer sourceFile.Close()

        //Leser innholdet til filen, en linje om gangen og deler linjene mellom 4 ulike deler.
       	/*scanner := bufio.NewScanner(sourceFile)
        for scanner.Scan() {
                line := scanner.Text()
                items := strings.Split(line, ";")

                fmt.Println(items[3])
        }*/

        //Lager filen der innholdet kan kopieres til.
        destinationFile, err := os.Create(dest)
        if err != nil {
                log.Fatal(err)
        }
        defer destinationFile.Close()

        //Tar input fra bruker og gj  r om kommandoer basert p   input.
        inputScanner := bufio.NewScanner(os.Stdin)
        fmt.Println("Type either \"convert\" to make a new file or \"average\" to get the average temperature.")
        inputScanner.Scan()
        input := inputScanner.Text()

        if (input == "convert" || input == "Convert") {
		inputScanner02 := bufio.NewScanner(os.Stdin)
                fmt.Println("Type either \"Celsius\" or \"Fahrenheit\" to chose output temperature.")
			inputScanner02.Scan()
			input02 := inputScanner02.Text()
			var fahrenheit float64;
				if (input02 == "Celsius" || input02 == "celsius" || input02 == "c" || input02 == "C") {
					scanner := bufio.NewScanner(sourceFile)
					for scanner.Scan() {
						line := scanner.Text()
						items := strings.Split(line, ";")

						fmt.Println(items)
					}

				}else if (input02 == "Fahrenheit" || input02 == "fahrenheit" || input02 == "f" || input02 == "F") {
					scanner02 := bufio.NewScanner(sourceFile)
					scanner02.Scan()

					for scanner02.Scan() {
						line02 := scanner02.Text()
						items02 := strings.Split(line02, ";")

						F, err := strconv.ParseFloat(items02[3], 64)
						if err != nil {
							log.Fatal(err)
						}
						fahrenheit = conv.CelsiusToFahrenheit(F)

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
				fmt.Println("Test \"C\" passed.")
			}else if (input == "Fahrenheit" || input == "fahrenheit" || input == "f" || input == "F") {
				fmt.Println("Test \"F\" passed.")
			}else {
				fmt.Println("Please type either \"Celsius \" or \"Fahrenheit\".")
			}

        }else {
                fmt.Println("Please type either \"convert\" or \"average\".")
        }
}

