package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"os"

	"github.com/cmaurer/go-lo/codegen/types"
)

var packageName = "slice"

func main() {
	// load the operations file
	cwd := os.Args[1]
	operationsBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/%s", cwd, "operations.json"))
	if err != nil {
		fmt.Println(err)
		return
	}

	var operations *types.Operations
	err = json.Unmarshal(operationsBytes, &operations)
	if err != nil {
		fmt.Println(err)
		return
	}

	outputDir := fmt.Sprintf("%s/generated", cwd)
	// remove any previously generated code
	err = os.RemoveAll(outputDir)
	if err != nil {
		fmt.Println(err)
		return
	}

	// ensure the generated dir exists
	err = os.Mkdir(outputDir, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, operation := range operations.Operations {

		templateFileName := fmt.Sprintf("%s/%s", cwd, operation.Template)
		outputFileName := fmt.Sprintf("%s/%s.go", outputDir, operation.Name)

		// create file
		f, err := os.Create(outputFileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		f.WriteString(fmt.Sprintf("package %s\n\n", packageName))

		templateFile, err := ioutil.ReadFile(templateFileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		tmpl := template.New(operation.Name)
		tmpl, err = tmpl.Parse(string(templateFile))
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, datatype := range operation.Datatypes {
			err = tmpl.Execute(f, datatype)
			if err != nil {
				fmt.Println(err)
				return
			}
		}

	}

}
