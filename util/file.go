// Package util is a set of utilitary things.
package util

import (
	"archive/zip"

	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Read file and return an array byte
func Read(filename string) []byte {
	// read in the contents of the localfile.data
	data, err := ioutil.ReadFile(filename)
	// if our program was unable to read the file
	// print out the reason why it can't
	if err != nil {
		fmt.Println(err)
	}

	// if it was successful in reading the file then
	// print out the contents as a string
	return data

}

// CreateTempFile make a temp folder
func CreateTempFile(prefix string) string {
	name, err := ioutil.TempDir(os.TempDir(), prefix)
	if err != nil {
		fmt.Println("Error creando el directorio temporal", err)
		return ""
	}
	return name
}

// Write function write a file
func Write(filename string, data []byte) {

	err := ioutil.WriteFile(filename, data, os.ModePerm)
	if err != nil {
		fmt.Println(err)
	}
}

// NewTempFile funtion create new file into system temp folder
func NewTempFile(fileName string) (f *os.File, err error) {
	return ioutil.TempFile(os.TempDir(), fileName)

}

// RenameFile permit rename a file with other name
func RenameFile(oldFileName, newFileName string) {
	err := os.Rename(oldFileName, newFileName)
	if err != nil {
		fmt.Println("No se pudo renombrar", err)
	}
}

// OpenDocument open de office document
func OpenDocument(src string) (a *zip.ReadCloser) {

	r, err := zip.OpenReader(src)
	if err != nil {
		fmt.Println(err)
	}
	return r

}

// Unzip file
// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, err
	}
	defer r.Close()

	for _, f := range r.File {

		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, err
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, err
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, err
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, err
		}
	}
	return filenames, nil
}

// Copy function is used for copy un file into other file
func Copy(src, dst string, BUFFERSIZE int64) error {
	sourceFileStat, err := os.Stat(src)
	if err != nil {
		return err
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file.", src)
	}

	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	_, err = os.Stat(dst)
	if err == nil {
		return fmt.Errorf("File %s already exists.", dst)
	}

	destination, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destination.Close()

	if err != nil {
		panic(err)
	}

	buf := make([]byte, BUFFERSIZE)
	for {
		n, err := source.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}
		if n == 0 {
			break
		}

		if _, err := destination.Write(buf[:n]); err != nil {
			return err
		}
	}
	return err
}
