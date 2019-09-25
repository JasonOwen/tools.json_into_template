package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"text/template"
)

func main() {

	var inputFile, templateFile, outputFile string

	flag.StringVar(&inputFile, "i", "", "Input JSON file")
	flag.StringVar(&templateFile, "t", "", "Template file")
	flag.StringVar(&outputFile, "o", "", "Output file")
	flag.Parse()

	t, err := template.ParseFiles(templateFile)
	if err != nil {
		fmt.Println("Error Reading Template File")
		panic(err)
	}

	b, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Error Reading JSON File")
		panic(err)
	}
	m := map[string]interface{}{}
	if err := json.Unmarshal(b, &m); err != nil {
		fmt.Println("Error Unmarshalling JSON File")
		panic(err)
	}

	if outputFile == "" {
		if err := t.Execute(os.Stdout, m); err != nil {
			fmt.Println("Error Executing Template File")
			panic(err)
		}
	} else {
		f, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("Error Creating Output File")
			panic(err)
		}
		if err := t.Execute(f, m); err != nil {
			fmt.Println("Error Executing Template File")
			panic(err)
		}
	}
}
