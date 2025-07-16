package main

import (
	"errors"
	"flag"
	
	"os"
)

type inputFile struct {
	filepath  string
	separator string
	pretty    bool
}

func main() {
	
	// fmt.Println(fileData,errors)
	

}
func getFileData() (inputFile, error) {
	if len(os.Args) < 2 {
		return inputFile{},errors.New("A filepath is required");




	}
	//Defining default values for separator and pretty value 
	//Preparing three args 1: paraname 2: default value 3: short desc 

	separator:=flag.String("separator","comma","column-separator")
	pretty :=flag.Bool("pretty",false ,"A pretty json conversion")

	flag.Parse()
	filepath:=flag.Arg(0)


	if (*separator!="comma" && *separator!="semicolon"){
		return inputFile{},errors.New("The separator value was not valid")
	}
  


   //if neither of the above things weren't caught then we send the input data 

	return inputFile{filepath,*separator,*pretty},nil




}