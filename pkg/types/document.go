package types

//EbmlDocument is the overarching document structure for an EBML doc
type EbmlDocument struct {
	Header   Header
	Segments []Segment
}

//Segment contains all the information about the individual EBML segments
type Segment struct {
	points    []Point
	tracks    []Track
	tags      []Tag
	files     []AttachedFile
	chapters  []Entry
	seekHeads []SeekHead
	infos     []Info
	clusters  []Cluster
}

//Header contains metadata about the document
type Header struct {
	EBMLVersion        uint
	EBMLReadVersion    uint
	EBMLMaxIDLength    uint
	EBMLMaxSizeLength  uint
	DocType            string
	DocTypeVersion     uint
	DocTypeReadVersion uint
}
