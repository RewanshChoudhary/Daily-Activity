package main

import (
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"path/filepath"

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

func checkValidFile( filename string ) (bool,error){
       fileExt:=filepath.Ext(filename)
	   if (fileExt!=".csv"){
		return false,fmt.Errorf("The extension was not csv instead it was : %v",fileExt)


	   }
          
	   if _,err:=os.Stat(filename);err!=nil && os.IsNotExist(err){
		return false,fmt.Errorf("The given file does not exist %v",filename)
	   }
	   return true ,nil



}


func processCSVFile(fileData inputFile ,writerChannel chan <- map[string]string){
	data,errors:=os.Open(fileData.filepath)

	if (errors!=nil){
		fmt.Errorf("Unexpected error 42069: %v",errors)

	}
	
	

	var headers [] string

	reader := csv.NewReader(data)
	if (fileData.separator=="semicolon"){
		reader.Comma=';'

	}
	
	headers,err:=reader.Read()
	check(err)

	for {
		line,err1:=reader.Read()


		if (err1==io.EOF){
			close(writerChannel)

			
		}else if err1!=nil {
		 exitGracefully(err1)

		}

		record,err:=processLine(headers,line)

		if err!=nil {
			fmt.Printf("Line :%sError: :%s line",line ,err)

		}

		writerChannel<-record



	}
  
}

func processLine(headers []string, dataLine []string) (map[string]string, error) {
 if (len(headers)!=len(dataLine)){
        return nil,errors.New("The number of headers and provided line does not match")


 }
recordMap:=make(map[string]string)

 for i,value :=range headers {
	recordMap[value]=dataLine[i]

 }

 return recordMap,nil


	

}

func exitGracefully(err1 error) {
	panic("Unexpected Error")
}

func check(er error ){
	if (er!=nil){
		exitGracefully(er)

	}
}