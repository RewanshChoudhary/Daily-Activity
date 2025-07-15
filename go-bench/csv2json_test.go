package main

import (
	"flag"
	"os"
	"testing"
)


func test_getFileType(t *testing.T){

	 tests:=[] struct{
		name string 
		want inputFile
		wantEr bool
		osArgs [] string 
		

	}{
		{"Default parameters:",inputFile{"index.csv","comma",false},false  ,[]string{"zsh","test.csv"}},
		{"No Parameter: ",inputFile{},true,[] string {}},
		{"Semi-colon",inputFile{"test.csv","semicolon",false},false,[] string {"zsh","--pretty","--separator=semicolon","test.csv"}},
        {"Pretty Enabled",inputFile{"test.csv","comma",true},false,[] string {"zsh","--pretty","test.csv"}},
		{"Pretty and Semicolon enabled",inputFile{"test.csv","semicolon",true},false,[]string {"zsh","--pretty","--separator=semicolon","test.csv"}},
		{"No separator ",inputFile{},true,[] string {"zsh","--separator=pipe"}},



	}

	for _,tt:=range tests{
		t.Run(tt.name,func (t *testing.T){
			actualArgs :=os.Args

			defer func (){
				os.Args=actualArgs
				flag.CommandLine=flag.NewFlagSet(tt.osArgs[0],flag.ContinueOnError)

			}()
			

		})
	}



	

}