package main

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"path/filepath"
	"strings"
	"os"
)

type inputFile struct {
	filepath  string
	separator string
	pretty    bool
}

func main() {

	flag.Usage = func() {
		fmt.Printf("Usage options %s [options]<cvsFile>\n Options \n", os.Args[0])
		flag.PrintDefaults()

	}
	fileData, err := getFileData()

	if err != nil {
		fmt.Print("1")
		exitGracefully(err)

	}
	if _, err := checkValidFile(fileData.filepath); err != nil {
		fmt.Print("2")
		exitGracefully(err)

	}
	writerChannel := make(chan map[string]string)

	done := make(chan bool)
	go processCSVFile(fileData, writerChannel)
	go writeJsonFile(fileData.filepath, writerChannel, done, fileData.pretty)
	<-done

}
func getFileData() (inputFile, error) {
	if len(os.Args) < 2 {
		return inputFile{}, errors.New("A filepath is required")

	}
	//Defining default values for separator and pretty value
	//Preparing three args 1: paraname 2: default value 3: short desc

	separator := flag.String("separator", "comma", "column-separator")
	pretty := flag.Bool("pretty", false, "A pretty json conversion")

	flag.Parse()
	filepath := flag.Arg(0)

	if *separator != "comma" && *separator != "semicolon" {
		return inputFile{}, errors.New("The separator value was not valid")
	}

	//if neither of the above things weren't caught then we send the input data

	return inputFile{filepath, *separator, *pretty}, nil

}

func checkValidFile(filename string) (bool, error) {
	fileExt := filepath.Ext(filename)
	if fileExt != ".csv" {
		return false, fmt.Errorf("The extension was not csv instead it was : %v", fileExt)

	}

	if _, err := os.Stat(filename); err != nil && os.IsNotExist(err) {
		return false, fmt.Errorf("The given file does not exist %v", filename)
	}
	return true, nil

}

func processCSVFile(fileData inputFile, writerChannel chan<- map[string]string) {
	data, err1 := os.Open(fileData.filepath)
if err1!= nil {
    exitGracefully(err1)
}
defer data.Close()

	var headers []string
	fmt.Print("The reader was instantiated")

	reader := csv.NewReader(data)
	if fileData.separator == "semicolon" {
		reader.Comma = ';'



	}
    
	headers, err := reader.Read()
	check(err)
	//fmt.Print("The headers were read")

	for {
		//fmt.Print("The lines will be read")

		line, err1 := reader.Read()

		// fmt.Print("The line was read")

		if err1 == io.EOF {
			close(writerChannel)
			break

		} else if err1 != nil {
			//fmt.Print("3")
			exitGracefully(err1)

		}

		record, err := processLine(headers, line)

		if err != nil {
			fmt.Printf("\n Line length mismatch!\nLine: %#v\nError: %v\n", line, err)


		}

		writerChannel <- record

	}

}

func processLine(headers []string, dataLine []string) (map[string]string, error) {
	if len(headers) != len(dataLine) {
		return nil, errors.New("The number of headers and provided line does not match")

	}
	recordMap := make(map[string]string)

	for i, value := range headers{
		recordMap[value] = dataLine[i]

	}
	fmt.Println("Parsed row:", recordMap)

	return recordMap, nil

}

func exitGracefully(err1 error) {
    fmt.Fprintf(os.Stderr, "Fatal error: %v\n", err1)
    os.Exit(1)
}

func check(er error) {
	if er != nil {
		fmt.Print("4")
		exitGracefully(er)

	}
}

func writeJsonFile(filePath string, writerChannel <-chan map[string]string, done chan<- bool, pretty bool) {
	writerString := getStringerWriter(filePath)
	jsonFunc, breakLine := getJson(pretty)

	fmt.Print("Writing the json file")
	writerString("["+breakLine, false)

	first := true
	for {
		record, more := <-writerChannel

		if more {
			if !first {
				writerString(","+breakLine, false)

			} else {
				first = false
				writerString(breakLine, false)

			}
			jsonData := jsonFunc(record)
			writerString(jsonData, false)

		} else {
			writerString("]"+breakLine, true)
			fmt.Print("Completed .....")
			done <- true

			break

		}

	}

}

func getStringerWriter(csvPath string) func(string, bool) {
	jsonDir := filepath.Dir(csvPath)
	jsonName := fmt.Sprintf("%s.json", strings.TrimSuffix(filepath.Base(csvPath), ".csv"))
	finalLocation := filepath.Join(jsonDir, jsonName)

	f, err := os.Create(finalLocation)
	check(err)

	return func(data string, close bool) {
		_, err = f.WriteString(data)

		if close {
			f.Close()
		}

	}

}

func getJson(pretty bool) (func(map[string]string) string, string) {

	var jsonFunc func(map[string]string) string
	var breakPoint string

	if pretty {
		breakPoint = "\n"

		jsonFunc = func(record map[string]string) string {
			jsonData, _ := json.MarshalIndent(record, "  ", "  ")
			return "   " + string(jsonData)

		}

	} else {
		breakPoint = ""

		jsonFunc = func(record map[string]string) string {
			jsonData, _ := json.Marshal(record)

			return string(jsonData)
		}

	}
	return jsonFunc, breakPoint
}
