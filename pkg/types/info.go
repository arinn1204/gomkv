package types

import "github.com/google/uuid"

//Info contains the general information about the segment
type Info struct {
	SegmentUID      *uuid.UUID               `json:",omitempty"`
	SegmentFilename string                   `json:",omitempty"`
	NextUID         *uuid.UUID               `json:",omitempty"`
	NextFilename    string                   `json:",omitempty"`
	PrevUID         *uuid.UUID               `json:",omitempty"`
	PrevFilename    string                   `json:",omitempty"`
	SegmentFamily   *uuid.UUID               `json:",omitempty"`
	TimestampScale  uint                     `json:",omitempty"`
	Duration        float32                  `json:",omitempty"`
	DateUTC         uint64                   `json:",omitempty"`
	Title           string                   `json:",omitempty"`
	MuxingApp       string                   `json:",omitempty"`
	WritingApp      string                   `json:",omitempty"`
	Translations    []TranslationInformation `json:",omitempty"`
}

//TranslationInformation is the mapping between this segment and other segments in a given chapter
type TranslationInformation struct {
	editionUID   uint   `json:",omitempty"`
	codec        uint   `json:",omitempty"`
	translateIDs []byte `json:",omitempty"`
}
