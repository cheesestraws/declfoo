package main

import "fmt"

// Constructor table
type SRDConstructor func (from []byte, entryLoc uint32, offset uint32) SRDatum

var srdConstructors map[string]SRDConstructor = map[string]SRDConstructor{
	"ignore": always_return(srIgnore{}),
	"cString": mksrString,
	"SExecBlock": mksrExecBlock,
}

type SRDatum interface {
	PrettyString() string
}

// ignore
type srIgnore struct{}
var _ = SRDatum(srIgnore{})

func (s srIgnore) PrettyString() string {
	return ""
}

// cString
type srString string
var _ = SRDatum(srString(""))

func (s srString) PrettyString() string {
	return "    \"" + string(s) + "\""
}

func mksrString(from []byte, entryLoc uint32, offset uint32) SRDatum {
	actualOffset := entryLoc + offset
	
	var bs []byte
	for from[actualOffset] != 0 {
		bs = append(bs, from[actualOffset])
		actualOffset++
	}
	
	return srString(bs)
}


// sExecBlock
type srExecBlock struct {
	offset uint32
	blockSize uint32
	cpuid uint8
	codeOffset uint32
	codeSize uint32
}

func (s srExecBlock) PrettyString() string {
	return fmt.Sprintf("    Code for 680%d0: %d bytes starting at 0x%06x", s.cpuid, s.codeSize, s.offset + 8 + s.codeOffset)
}

func mksrExecBlock(from []byte, entryLoc uint32, offset uint32) SRDatum {
	actualOffset := entryLoc + offset
	blockSize := (uint32(from[actualOffset]) << 24) | 
		(uint32(from[actualOffset + 1]) << 16) |
		(uint32(from[actualOffset + 2]) << 8) | 
		(uint32(from[actualOffset + 3]))
		
	cpuid := from[actualOffset+5]
	
	codeOffset := (uint32(from[actualOffset + 8]) << 24) | 
		(uint32(from[actualOffset + 9]) << 16) |
		(uint32(from[actualOffset + 10]) << 8) | 
		(uint32(from[actualOffset + 11]))



	return srExecBlock{
		offset: actualOffset,
		blockSize: blockSize,
		cpuid: cpuid,
		codeOffset: codeOffset,
		codeSize: blockSize - 8 - codeOffset,
	}
}


// utility

func always_return(val SRDatum) SRDConstructor {
	return func (from []byte, entryLoc uint32, offset uint32) SRDatum {
		return val
	}
}
