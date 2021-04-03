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

var fieldTypes map[SResourceFieldID]string = map[SResourceFieldID]string{
	// General fields

	1: "ignore",
	2: "cString",
	3: "ignore",
	4: "ignore",
	5: "ignore",
	6: "ignore",
	7: "ignore",
	8: "ignore",
	10: "ignore",
	11: "ignore",
	12: "ignore",
	13: "ignore",
	15: "ignore",
	16: "ignore",
	17: "ignore",
	108: "ignore",
	
	// Board sResource fields
	
	32: "ignore",
	33: "ignore",
	34: "SExecBlock",
	35: "ignore",
	36: "ignore",
	38: "SExecBlock",
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
	decodedData map[SResourceFieldID]SRDatum
}

func (s *SResource) populate(from []byte) {
	s.fieldEntryLoc = make(map[SResourceFieldID]uint32)
	s.fieldData = make(map[SResourceFieldID]uint32)
	s.decodedData = make(map[SResourceFieldID]SRDatum)

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
		s.decodedData[fieldID] = decodesrResourceField(fieldID, from, entryOffset, data)
		
		entryOffset += 4
	}
}

func decodesrResourceField(fid SResourceFieldID, from []byte, entryLoc uint32, offset uint32) SRDatum {
	// find the field type
	ft, ok := fieldTypes[fid]
	if !ok {
		ft = "ignore"
	}

	// find the function to decode that
	fn := srdConstructors[ft]
	
	// do the decoding
	return fn(from, entryLoc, offset)
}

func (s SResource) PrettyString() string {
	var str string
	for _, fid := range s.fields {
		str += fmt.Sprintf("  %s => 0x%06x\n", fid.PrettyString(), s.fieldData[fid]);
		fldstr := s.decodedData[fid].PrettyString()
		if fldstr != "" {
			str += fldstr + "\n"
		}
	}
	
	return str
}