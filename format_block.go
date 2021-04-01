package main

import "errors"

type FormatBlock struct {
	byteLanes byte
	format byte
	directoryOffset uint32
	directoryAddress uint32
}

var (
	ErrNoTestPattern = errors.New("could not find test pattern; invalid ROM?")
	ErrNotAppleFormatROM = errors.New("not an Apple format ROM")
)

// For the time being we assume the format block is at the end of the ROM
// and that the end of the ROM is the top of nubus space.
func ExtractFormatBlock(bs []byte) (*FormatBlock, error) {
	var fb FormatBlock
	l := len(bs)
	
	// Check the test pattern
	if !(bs[l-6] == 0x5a && bs[l-5] == 0x93 && bs[l-4] == 0x2b && bs[l-3] == 0xc7) {
		return nil, ErrNoTestPattern
	}
	
	// Check this is an Apple format ROM
	if bs[l-7] != 0x01 {
		return nil, ErrNotAppleFormatROM
	}
	
	fb.byteLanes = bs[l-1]
	fb.format = bs[l-7]
	
	// find the offset
	var offset uint32
	offset = 0xff000000
	offset |= uint32(bs[l-17]) | (uint32(bs[l-18]) << 8) | (uint32(bs[l-19]) << 16)
	offset = (0xffffffff ^ offset) + 1
	fb.directoryOffset = offset
	
	// find the actual address of the directory in the ROM
	var addr uint32
	addr = (uint32(l)-20) - offset
	fb.directoryAddress = addr	
	
	return &fb, nil
}