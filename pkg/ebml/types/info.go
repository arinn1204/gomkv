package types

import "github.com/google/uuid"

//Info contains the general information about the segment
type Info struct {
	current       SegmentInformation
	previous      SegmentInformation
	next          SegmentInformation
	segmentFamily uuid.UUID
	timescale     uint
	duration      float32
	date          uint64
	title         string
	muxingApp     string
	writingApp    string
	translations  []TranslationInformation
}

//SegmentInformation is a reference to the location in the stream that the filename is
type SegmentInformation struct {
	UID      uuid.UUID
	filename string
}

//TranslationInformation is the mapping between this segment and other segments in a given chapter
type TranslationInformation struct {
	editionUID   uint
	codec        uint
	translateIDs []byte
}
