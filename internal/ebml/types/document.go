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

type Document struct {
	header   Header
	segments []Segment
}

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

type Header struct {
	maxIDLength   int
	maxSizeLength int
	docType       string
	version       int
}
