package main

import "fmt"

type ROM struct {
	formatBlock *FormatBlock
}

func DissectROM(bs []byte) (*ROM, error) {
	var rom ROM
	var err error
	
	rom.formatBlock, err = ExtractFormatBlock(bs)
	if err != nil {
		return nil, err
	}
	
	return &rom, nil
}

func (r *ROM) PrettyString() string {
	return "farts"
}

func (r *ROM) Dump() {
	fmt.Println(r.PrettyString())
}