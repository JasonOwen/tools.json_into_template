package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
	"text/template"
)

func findString(regex, str string) string {
	re := regexp.MustCompile(regex)
	return re.FindString(str)
}

func matchString(str, matchstring string) bool {
	if matchstring == str {
		return true
	}
	return false
}

func regexReplace(str, matchstring, replacewith string) string {
	var re = regexp.MustCompile(matchstring)

	return re.ReplaceAllString(str, replacewith)
}

func main() {

	var inputFile, templateFile, outputFile, feedInVariables, unknownVar string

	flag.StringVar(&inputFile, "i", "", "Input JSON file")
	flag.StringVar(&templateFile, "t", "", "Template file")
	flag.StringVar(&outputFile, "o", "", "Output file")
	flag.StringVar(&feedInVariables, "v", "", "Feed In Variables")
	flag.StringVar(&unknownVar, "u", "Unknown", "No data string replacement")

	flag.Parse()
	templateName := path.Base(templateFile)

	t, err := template.New(templateName).Option(fmt.Sprintf("missingkey=%s", unknownVar)).Funcs(template.FuncMap{
		"findString":   findString,
		"matchString":  matchString,
		"regexReplace": regexReplace,
	}).ParseFiles(templateFile)

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
	m = addFeedInVariables(feedInVariables, m)

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

func addFeedInVariables(feedInCSV string, m map[string]interface{}) map[string]interface{} {
	for _, feedIn := range strings.Split(feedInCSV, ",") {
		if strings.Contains(feedIn, "=") {
			parameterSet := strings.Split(feedIn, "=")

			m[parameterSet[0]] = parameterSet[1]
		}
	}
	return m
}
