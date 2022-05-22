package types

import (
	"github.com/arinn1204/gomkv/internal/ebml/types/attachment"
)

type Chapter struct{}
type Cluster struct{}
type Cue struct{}
type Info struct{}
type SeekHead struct{}
type Tag struct{}
type Track struct{}

type Segment struct {
	cue        Cue
	tracks     []Track
	tags       []Tag
	attachment attachment.Attachment
	chapter    Chapter
	seekHeads  []SeekHead
	infos      []Info
	clusters   []Cluster
}

type Header struct {
	maxIDLength   int
	maxSizeLength int
	docType       string
	version       int
}

type Document struct {
	header   Header
	segments []Segment
}
