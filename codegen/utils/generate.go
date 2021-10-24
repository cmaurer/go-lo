package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"text/template"

	"github.com/cmaurer/go-lo/codegen/types"
)

var packageName = "utils"

func main() {
	// load the operations file
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		panic("No caller information")
	}
	cwd := path.Dir(filename)

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

	generateCode(cwd, outputDir, operations)
	generateTests(cwd, outputDir, operations)
	generateBenchmarks(cwd, outputDir, operations)

}

func generateCode(cwd string, outputDir string, operations *types.Operations) {

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

		_, err = f.WriteString(fmt.Sprintf("package %s\n\n", packageName))
		if err != nil {
			fmt.Println(err)
			return
		}

		if len(operation.Imports) > 0 {
			_, err = f.WriteString("import (\n\n")
			if err != nil {
				fmt.Println(err)
				return
			}
			for _, imprt := range operation.Imports {
				_, err = f.WriteString(fmt.Sprintf("\"%s\"\n", imprt))
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			_, err = f.WriteString(")\n\n")
			if err != nil {
				fmt.Println(err)
				return
			}
		}

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

func generateTests(cwd string, outputDir string, operations *types.Operations) {

	for _, operation := range operations.Operations {

		templateFileName := fmt.Sprintf("%s/%s", cwd, operation.TestTemplate)
		outputFileName := fmt.Sprintf("%s/%s_test.go", outputDir, operation.Name)

		// create file
		f, err := os.Create(outputFileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		_, err = f.WriteString(fmt.Sprintf(`package %s

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
		`, packageName))
		if err != nil {
			fmt.Println(err)
			return
		}

		templateFile, err := ioutil.ReadFile(templateFileName)
		if err != nil {
			fmt.Println(err)
			return
		}

		tmpl := template.New(operation.Name).Funcs(template.FuncMap{
			"dataTypeGenerator": dataTypeGenerator,
		})
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

func generateBenchmarks(cwd string, outputDir string, operations *types.Operations) {

	for _, operation := range operations.Operations {

		if operation.BenchTemplate == "" {
			continue
		}
		templateFileName := fmt.Sprintf("%s/%s", cwd, operation.BenchTemplate)
		outputFileName := fmt.Sprintf("%s/%s_bench_test.go", outputDir, operation.Name)

		// create file
		f, err := os.Create(outputFileName)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		_, err = f.WriteString(fmt.Sprintf(`package %s

import (
	"testing"
)
		`, packageName))
		if err != nil {
			fmt.Println(err)
			return
		}

		templateFile, err := ioutil.ReadFile(templateFileName)
		if err != nil {
			fmt.Println(err)
			return
		}

		tmpl := template.New(operation.Name).Funcs(template.FuncMap{
			"dataTypeGenerator": dataTypeGenerator,
		})
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

func dataTypeGenerator(datatype string, value int) string {
	switch datatype {
	case "string":
		return fmt.Sprintf("\"test%d\"", value)
	case "int":
		return fmt.Sprintf("%d", value)
	case "int8":
		return fmt.Sprintf("%d", value)
	case "int16":
		return fmt.Sprintf("%d", value)
	case "int32":
		return fmt.Sprintf("%d", value)
	case "int64":
		return fmt.Sprintf("%d", value)
	case "uint":
		return fmt.Sprintf("%d", value)
	case "uint8":
		return fmt.Sprintf("%d", value)
	case "uint16":
		return fmt.Sprintf("%d", value)
	case "uint32":
		return fmt.Sprintf("%d", value)
	case "uint64":
		return fmt.Sprintf("%d", value)
	case "float32":
		return fmt.Sprintf("%d.%d", value, value)
	case "float64":
		return fmt.Sprintf("%d.%d", value, value)
	case "complex64":
		var a float64 = (2.4 * float64(value))
		var b float64 = 3.14
		return fmt.Sprintf("%g", complex64(complex(a, b)))
	case "complex128":
		var a float64 = (2.4 * float64(value))
		var b float64 = 3.14
		return fmt.Sprintf("%g", complex128(complex(a, b)))
	}
	return "nil"
}
