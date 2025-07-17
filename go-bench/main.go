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
			break


			
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

func writeJsonFile(filePath string ,writerChannel <- chan map[string]string ,done chan <- bool,pretty bool){
	writerString :=getStringerWriter(filePath)
	jsonFunc,breakLine:=getJson(pretty)

	fmt.Print("Writing the json file")
	writerString("["+breakLine,false)

	first:=true
	for {
		record,more:=<-writerChannel

		if more{
			if !first{
				writerString("["+breakLine,false)

			}else {
				first=false

			}
			jsonData:=jsonFunc(record)
			writerString(jsonData,false)

		}else {
			writerString("]"+breakLine,true)
			fmt.Print("Completed .....")
			done <-true

			break;

		}


	}
	
}

func getStringerWriter(csvPath string) func (string,bool) {
	jsonDir:=filepath.Dir(csvPath)
	jsonName:=fmt.Sprintf("%s.json",strings.TrimSuffix(filepath.Base(csvPath),".csv"))
	finalLocation:=filepath.Join(jsonDir,jsonName)

	f,err:=os.Create(finalLocation)
	check(err)

	return func (data string ,close bool) {
		_,err =f.WriteString(data)
		

		if (close){
			f.Close()
		}


	}




}

func getJson(pretty bool ) (func (map[string ]string  )string ,string ){

	var jsonFunc func(map[string]string) string
	var breakPoint string

	if (pretty){
		breakPoint="\n"

		jsonFunc=func(record map[string ]string )string {
			jsonData,_:=json.MarshalIndent(record,"  ","  ")
			return "   "+string(jsonData)



		}

	}else {
		breakPoint=""

		jsonFunc=func (record map[string ]string )string {
			jsonData,_:=json.Marshal(record)

			return string(jsonData)
		}



	}
	return jsonFunc,breakPoint
}