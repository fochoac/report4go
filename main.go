package main

import (
	"archive/zip"
	"encoding/xml"
	"fmt"

	file "github.com/fochoac/report4go/util"
)

func main() {
	const docxTemplate = "empleados.docx"
	tempFolderPath := file.CreateTempFile("prueba")
	tempTemplate := tempFolderPath + "/empleados.docx"
	//	tempZipTemplate := tempFolderPath + "/empleados.zip"

	var buffer int64 = 1024
	error := file.Copy(docxTemplate, tempTemplate, buffer)
	if error != nil {
		fmt.Errorf("error: %s ", error)
	}
	//file.RenameFile(tempTemplate, tempZipTemplate)
	stream := file.OpenDocument(tempTemplate)
	defer stream.Close()
	documento := getDocument(stream)

	/*	rutas, errorUnzip := file.Unzip(tempTemplate, tempFolderPath+"/empleados")
		if errorUnzip != nil {
			fmt.Println("Error al descomprimir", errorUnzip)
		}

		for index := 0; index < len(rutas); index++ {
			fmt.Println(rutas[index])
		}*/
	a, err := documento.Open()
	if err != nil {
		panic(err)
	}

	decoder := xml.NewDecoder(a)

	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		//	fmt.Println(token)

		switch element := token.(type) {
		case xml.StartElement:
			{
				//	fmt.Println(element.Name)
				/*	for _, attr := range element.Attr {
					/*if attr.Name.Local == "Ignorable" {
						fmt.Println(attr.Name.Local)
					} else {*/
				//	attr.Name.Local
				//	fmt.Printf(" %s", attr.Value)
				//doc.Scheme[attr.Name.Local] = attr.Value
				//}
				//	}

				if element.Name.Local == "body" {
					decoderElement(decoder)
				}
			}
		}

	}
}
func decoderElement(decoder *xml.Decoder) {
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch element := token.(type) {
		case xml.StartElement:
			{
				printAttr(element)
				xml.NewTokenDecoder(token)
				decoderElement(decoder)

			}
		case xml.EndElement:
			{
				if element.Name.Local == "body" {
					break
				}
			}
		}
	}
}
func getDocument(read *zip.ReadCloser) (f *zip.File) {
	for _, f := range read.File {
		if f.Name == "word/document.xml" {
			return f
		}
	}
	return nil
}

func printAttr(element xml.StartElement) {
	for _, attr := range element.Attr {
		fmt.Println(attr.Name.Local)
	}
}
