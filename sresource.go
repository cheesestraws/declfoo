package main

import (
	"fmt"
)

type SResourceFieldID uint8

var fieldNames map[SResourceFieldID]string = map[SResourceFieldID]string{
	// General fields

	1: "sRsrcType",
	2: "sRsrcName",
	3: "sRsrcIcon",
	4: "sRsrcDrvrDir",
	5: "sRsrcLoadRec",
	6: "sRsrcBootRec",
	7: "sRsrcFlags",
	8: "sRsrcHWDevId",
	10: "MinorBaseOS",
	11: "MinorLength",
	12: "MajorBaseOS",
	13: "MajorLength",
	15: "sRsrcCicn",
	16: "sRsrcIcl8",
	17: "sRsrcIcl4",
	108: "sMemory",
	
	// Board sResource fields
	
	32: "BoardId",
	33: "PRAMInitData",
	34: "PrimaryInit",
	35: "STTimeOut",
	36: "VendorInfo",
	38: "SecondaryInit",
}

func (f SResourceFieldID) PrettyString() string {
	if _, ok := fieldNames[f]; ok {
		return fieldNames[f]
	}
	return fmt.Sprintf("field 0x%02x", uint8(f))
}


type SResource struct {
	ID uint8
	Offset uint32
	
	fields []SResourceFieldID
	fieldEntryLoc  map[SResourceFieldID]uint32
	fieldData map[SResourceFieldID]uint32
}

func (s *SResource) populate(from []byte) {
	s.fieldEntryLoc = make(map[SResourceFieldID]uint32)
	s.fieldData = make(map[SResourceFieldID]uint32)

	entryOffset := s.Offset
	
	// Each entry is an ID and an offset FROM THE TABLE ENTRY
	// The end-of-list marker is an ID of 0xff
	for from[entryOffset] != 0xFF {
		var data uint32
		
		fieldID := SResourceFieldID(from[entryOffset])
		
		data = (uint32(from[entryOffset + 1]) << 16) |
			(uint32(from[entryOffset + 2]) << 8) |
			(uint32(from[entryOffset + 3]))
						
		s.fields = append(s.fields, fieldID)
		s.fieldData[fieldID] = data
		s.fieldEntryLoc[fieldID] = entryOffset		
		
		entryOffset += 4
	}
}

func (s SResource) PrettyString() string {
	var str string
	for _, fid := range s.fields {
		str += fmt.Sprintf("  %s => 0x%06x\n", fid.PrettyString(), s.fieldData[fid]);
	}
	
	return str
}