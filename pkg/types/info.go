package types

import "github.com/google/uuid"

//Info contains the general information about the segment
type Info struct {
	SegmentUID      uuid.UUID
	SegmentFilename string
	NextUID         uuid.UUID
	NextFilename    string
	PrevUID         uuid.UUID
	PrevFilename    string
	SegmentFamily   uuid.UUID
	TimestampScale  uint
	Duration        float32
	DateUTC         uint64
	Title           string
	MuxingApp       string
	WritingApp      string
	Translations    []TranslationInformation
}

//TranslationInformation is the mapping between this segment and other segments in a given chapter
type TranslationInformation struct {
	editionUID   uint
	codec        uint
	translateIDs []byte
}
