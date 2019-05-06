package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	doc "github.com/fochoac/report4go/util"
)

func main() {
	var variables = []string{"${fechaCorte}"}
	fmt.Println("hola")
	xmlFile := getData()
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	var document doc.Document

	xml.Unmarshal(byteValue, &document)

	for _, P := range *document.Body.P {
		row := P.R

		for _, cadena := range variables {
			var texto *string
			if row.T != nil {
				texto = row.T
			} else {
				continue
			}

			if strings.Contains(*texto, cadena) {
				replaceString(texto, cadena)

			}
		}
	}

	printDocument(&document)
	resultado, err := xml.MarshalIndent(&document, "d", "")
	if err != nil {
		panic(err)
	}

	ioutil.WriteFile("salida.xml", resultado, os.ModePerm)

}
func replaceString(source *string, value string) {
	replace := strings.Replace(*source, value, "se cambio la vaina", -1)
	*source = replace

}
func printDocument(document *doc.Document) {
	for _, item := range *document.Body.P {
		if item.R.T != nil {
			fmt.Println(*item.R.T)
		}

	}
}
func getData() *os.File {
	xmlFile, err := os.Open("document.xml")

	if err != nil {
		panic(err)
	}
	return xmlFile
}
