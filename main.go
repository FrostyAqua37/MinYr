package main

import (
        "os"
        "log"
//      "io/ioutil"
        "bufio"
        "fmt"
        "strings"
        "strconv"
	"reflect"
	"errors"

	"github.com/FrostyAqua37/funtemps/conv"
)

func main() {

        src := "kjevik-temp-celsius-20220318-20230318.csv" //Sourcefile is assigned to "src".
  	dest := "kjevik-temp-fahr-20220318-20230318.csv" //Destinationfile is assigned to "dest".


        sourceFile, err := os.Open(src) //Opens the sourcfile.
        if err != nil { //Checks if there is an error or not. 
                log.Fatal(err)
        }
        defer sourceFile.Close() //Command to close the sourcfile last. 

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
					var lineCount int 

					for scanner.Scan() { //Loops through the file line by line and prints it out. 
						line := scanner.Text()
						items := strings.Split(line, ";") //Divides each line when a ";" appears in the line. 
						lineCount++

						if lineCount == 16756 {
							lastLine := "Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Eivind Chen."
							fmt.Println(lastLine)
							break
						}
						items02 := fmt.Sprint(items)
						fmt.Println(items02)
					}

				}else if (input02 == "Fahrenheit" || input02 == "fahrenheit" || input02 == "f" || input02 == "F") { //Checks if the input is one of the options listed. 
					scanner02 := bufio.NewScanner(sourceFile)
					scanner02.Scan()
					var fahrenheit float64
					var lineCount int 

					lastLine := fmt.Sprintln("Data er basert på gyldig data (per 18.03.2023) (CC BY 4.0) fra Meteorologisk institutt (MET);endringen er gjort av Eivind Chen.")

					if _, err := os.Stat(dest); err == nil {
					inputScanner := bufio.NewScanner(os.Stdin)
					fmt.Println("The files does exist, do you want to make a new file(Yes or No)?")
					inputScanner.Scan()
					input := inputScanner.Text()

						if input == "Yes" || input == "yes" {
							destinationFile, err := os.Create(dest)

							if err != nil {
               							log.Fatal(err)
       							}

								for scanner02.Scan() {
                                    line := scanner02.Text()
                                    items := strings.Split(line, ";")
                                    lineCount++

 			                        if lineCount == 16755 {
                                        break
                                    }

                                    f, err := strconv.ParseFloat(items[3], 64)
                                    	if err != nil {
                                            log.Fatal(err)
                                        }

				                    fahrenheit = conv.CelsiusToFahrenheit(f)

				                    fahrenheitString := fmt.Sprintf("%2.f", fahrenheit)
                                	items[3] = fahrenheitString

				                    items02 := fmt.Sprintln(items)
                                	destinationFile.Write([]byte(items02))
								}

                                if lineCount == 16755 {
                                    destinationFile.Write([]byte(lastLine))
                                }

                                fmt.Println("The temperature has now been converted from \"Celsius\" to \"Fahrenheit\".")
                                defer destinationFile.Close()


						} else if input == "No" || input == "no" {
								fmt.Println("File was not overwritten.")
						}

					} else if errors.Is(err, os.ErrNotExist) {
						
						destinationFile, err := os.Create(dest)
							if err != nil {
               					log.Fatal(err)
       						}
					
						for scanner02.Scan() {
							line := scanner02.Text()
							items := strings.Split(line, ";")
							lineCount++

							if lineCount == 16755 {
								break
							}

							f, err := strconv.ParseFloat(items[3], 64)
								if err != nil {
									log.Fatal(err)
								}
							
							fahrenheit = conv.CelsiusToFahrenheit(f)
							
							fahrenheitString := fmt.Sprintf("%1.f", fahrenheit)
							items[3] = fahrenheitString

							items02 := fmt.Sprintln(items)
							destinationFile.Write([]byte(items02))
						}

						if lineCount == 16755 {
							destinationFile.Write([]byte(lastLine))
						}

						fmt.Println("Temperaturverdiene har nå blitt konvertert fra Celsius til Fahrenheit.")
						defer destinationFile.Close()
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
				var lineCount int64 = 1
				var total int64

				for fileScanner.Scan() {
					line := fileScanner.Text()
					items := strings.Split(line, ";")
					lineCount++

					c, err := strconv.ParseInt(items[3], 10, 64)
					if reflect.Kind(c) != reflect.Int {
					continue
					}

					if err != nil {
						log.Fatal(err)
						}

					total += c
				}


				fmt.Println("All temperature values added together:", total)
				fmt.Println("Number of lines in file:", lineCount)
				fmt.Println("Average temperature(Celsius):", total/lineCount)



			}else if (input == "Fahrenheit" || input == "fahrenheit" || input == "f" || input == "F") {
				fileScanner := bufio.NewScanner(sourceFile)
				fileScanner.Scan()
				var lineCount float64 = 1
				var total float64
				var fahrenheit float64

				for fileScanner.Scan() {
					line := fileScanner.Text()
					items := strings.Split(line, ";")
					lineCount++

					f, err := strconv.ParseFloat(items[3], 64)
					fahrenheit = conv.CelsiusToFahrenheit(f)

					if reflect.Kind(f) != reflect.Int {
					continue
					}

					if err != nil {
                        log.Fatal(err)
                    }
//                                        fmt.Println(fahrenheit)

					total += fahrenheit
				}

				fmt.Println("All temperature values added together:", total)
				fmt.Println("Number of lines in file:", lineCount)
				fmt.Println("Average temperature(Fahrenheit):", total/lineCount)

			}else {
				fmt.Println("Please type either \"Celsius \" or \"Fahrenheit\".")
			}

        }else {
                fmt.Println("Please type either \"convert\" or \"average\".")
        }
	}


