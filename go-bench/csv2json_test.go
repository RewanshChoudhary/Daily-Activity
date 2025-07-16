package main

import (
	"flag"
	"os"
	"reflect"
	"testing"
)


func Test_getFileType(t *testing.T){

	 tests:=[] struct{
		name string 
		want inputFile
		wantEr bool
		osArgs [] string 
		

	}{
		{"Default parameters:",inputFile{"test.csv","comma",false},false  ,[]string{"zsh","test.csv"}},
		{"No Parameter: ",inputFile{},true,[] string {"zsh"}},
		{"Semi-colon",inputFile{"test.csv","semicolon",false},false,[] string {"zsh","--separator=semicolon","test.csv"}},
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

		os.Args=tt.osArgs
		got,err:=getFileData()

		if (err!=nil) != tt.wantEr{
			t.Errorf("getFileData error= %v and the want error %v",err,tt.wantEr);
			
		}
			if !reflect.DeepEqual(got,tt.want){
				t.Errorf("getFileData error =%v and the wanted result  %v",got ,tt.want)
			}

		})
	}


	


	

}