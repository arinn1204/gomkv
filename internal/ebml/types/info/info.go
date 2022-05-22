package info

import "github.com/google/uuid"

//Info contains the general information about the segment
type Info struct {
	current       Segment
	previous      Segment
	next          Segment
	segmentFamily uuid.UUID
	timescale     uint
	duration      float32
	date          uint64
	title         string
	muxingApp     string
	writingApp    string
	translations  []Translate
}

//Segment is a reference to the location in the stream that the filename is
type Segment struct {
	UID      uuid.UUID
	filename string
}

//Translate is the mapping between this segment and other segments in a given chapter
type Translate struct {
	editionUID   uint
	codec        uint
	translateIDs []byte
}
