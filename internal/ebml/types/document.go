package types

import (
	"github.com/arinn1204/gomkv/internal/ebml/types/attachment"
	"github.com/arinn1204/gomkv/internal/ebml/types/chapter"
	"github.com/arinn1204/gomkv/internal/ebml/types/cluster"
	"github.com/arinn1204/gomkv/internal/ebml/types/cue"
	"github.com/arinn1204/gomkv/internal/ebml/types/info"
)

type SeekHead struct{}
type Tag struct{}
type Track struct{}

type Document struct {
	header   Header
	segments []Segment
}

type Segment struct {
	cue        cue.Cue
	tracks     []Track
	tags       []Tag
	attachment attachment.Attachment
	chapter    chapter.Chapter
	seekHeads  []SeekHead
	infos      []info.Info
	clusters   []cluster.Cluster
}

type Header struct {
	maxIDLength   int
	maxSizeLength int
	docType       string
	version       int
}
