package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	srcFile, err := os.Open("source.txt")

	if err != nil {
		fmt.Println("Error opening the source file:", err)
		os.Exit(1)
	}
	defer srcFile.Close()

	dstFile, err := os.Create("destination.txt")
	if err != nil {
		fmt.Println("Error creating destination file:", err)
		os.Exit(1)
	}
	defer dstFile.Close()

	written, err := io.Copy(os.Stdout, srcFile)
	if err != nil {
		fmt.Println("Error copying file:", err)
		os.Exit(1)
	}
	fmt.Println("copy successfull. Bytes writeen: \n", written)
	// io.copy function return (no of bytes copied , error if exists)
	// in this example we've learned the implementations about os.Open() , os.Create() ,io.copy() functions , filename.Close() also
}
