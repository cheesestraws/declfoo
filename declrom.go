package main

import "fmt"

type ROM struct {

}

func DissectROM(bs []byte) (*ROM, error) {
	var rom ROM
	
	return &rom, nil
}

func (r *ROM) PrettyString() string {
	return "farts"
}

func (r *ROM) Dump() {
	fmt.Println(r.PrettyString())
}