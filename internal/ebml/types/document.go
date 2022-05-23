package types

import (
	"github.com/arinn1204/gomkv/internal/ebml/types/attachment"
	"github.com/arinn1204/gomkv/internal/ebml/types/chapter"
	"github.com/arinn1204/gomkv/internal/ebml/types/cluster"
	"github.com/arinn1204/gomkv/internal/ebml/types/cue"
	"github.com/arinn1204/gomkv/internal/ebml/types/info"
	"github.com/arinn1204/gomkv/internal/ebml/types/seekhead"
	"github.com/arinn1204/gomkv/internal/ebml/types/tag"
	"github.com/arinn1204/gomkv/internal/ebml/types/track"
)

//EbmlDocument is the overarching document structure for an EBML doc
type EbmlDocument struct {
	header   Header
	segments []Segment
}

//Segment contains all the information about the individual EBML segments
type Segment struct {
	points    []cue.Point
	tracks    []track.Track
	tags      []tag.Tag
	files     []attachment.AttachedFile
	chapters  []chapter.Entry
	seekHeads []seekhead.SeekHead
	infos     []info.Info
	clusters  []cluster.Cluster
}

//Header contains metadata about the document
type Header struct {
	maxIDLength   int
	maxSizeLength int
	docType       string
	version       int
}
