package main

import "fmt"

type ROM struct {
	formatBlock *FormatBlock
	sResourceDirectory *SResourceDirectory
}

func DissectROM(bs []byte) (*ROM, error) {
	var rom ROM
	var err error
	
	rom.formatBlock, err = ExtractFormatBlock(bs)
	if err != nil {
		return nil, err
	}
	
	rom.sResourceDirectory, err = ExtractSResourceDirectory(bs, rom.formatBlock.directoryAddress)
	if err != nil {
		return nil, err
	}
	
	return &rom, nil
}

func (r *ROM) PrettyString() string {
	return r.sResourceDirectory.PrettyString()
}

func (r *ROM) Dump() {
	fmt.Println(r.PrettyString())
}