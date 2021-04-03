package main

import (
	"fmt"
)

type SResourceDirectory struct {
	ids []uint8
	dir map[uint8]SResource
}

func (s *SResourceDirectory) PrettyString() string {
	var str string
	for _, id := range s.ids {
		str += fmt.Sprintf("id 0x%02x", id)
		str += fmt.Sprintf(" @ 0x%06x\n", s.dir[id].Offset)
		
		str += s.dir[id].PrettyString()
		
		str += "\n"
	}
	
	return str
}

func ExtractSResourceDirectory(from []byte, offset uint32) (*SResourceDirectory, error) {
	// We should have some error checking here.  declare the type thus in
	// the hope it will guilt trip us into implementing that later

	// make our storage
	dir := SResourceDirectory{
		dir: make(map[uint8]SResource),
	}
	
	entryOffset := offset
	
	// Each entry is an ID and an offset FROM THE TABLE ENTRY
	// The end-of-list marker is an ID of 0xff
	for from[entryOffset] != 0xFF {
		var actualOffset uint32
		
		actualOffset = (uint32(from[entryOffset + 1]) << 16) |
			(uint32(from[entryOffset + 2]) << 8) |
			(uint32(from[entryOffset + 3]))
		
		actualOffset += entryOffset
		
		dir.ids = append(dir.ids, from[entryOffset])
		
		sr := SResource{
			ID: from[entryOffset], 
			Offset: actualOffset,
		}
		sr.populate(from)
		
		dir.dir[from[entryOffset]] = sr

		
		entryOffset += 4
	}
	
	return &dir, nil
}