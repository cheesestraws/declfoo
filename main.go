package main

import (
	"os"
	"fmt"
	"io/ioutil"
)

func usage() {
	fmt.Println("usage: declfoo [filename]")
}

func main() {
	// Args pls
	if len(os.Args) != 2 {
		usage()
		return
	}
	
	filename := os.Args[1]
	
	// Load the file
	bytes, err := ioutil.ReadFile(filename)
	
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}	
	
	rom, err := DissectROM(bytes)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}	
	
	rom.Dump()
}