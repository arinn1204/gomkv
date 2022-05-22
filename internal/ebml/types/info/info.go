package info

import "github.com/google/uuid"

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

type Segment struct {
	uid      uuid.UUID
	filename string
}

type Translate struct {
	editionUid   uint
	codec        uint
	translateIds []byte
}
