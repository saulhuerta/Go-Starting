package main

import (
	"fmt"
	"os"
)

func main(){

	myFile := createFile("test.txt")
	defer closeFile(myFile)
	writeFile(myFile)
}

func createFile(fileNameParam string) *os.File{

	fmt.Println("Creating file...")
	f, err := os.Create(fileNameParam)

	if err != nil {
		panic(err)
	}

	return f

}

func writeFile(fileNameParam *os.File)  {

	fmt.Println("Writing...")
	fmt.Fprintln(fileNameParam, "Info")

}

func closeFile(f *os.File)  {

	fmt.Fprintln(f, "Closing...")
	err := f.Close()

	if err != nil {
		fmt.Fprintln(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

}