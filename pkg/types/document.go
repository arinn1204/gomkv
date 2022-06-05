package types

//EbmlDocument is the overarching document structure for an EBML doc
type EbmlDocument struct {
	Header   *Header
	Segments []Segment
}

//Segment contains all the information about the individual EBML segments
type Segment struct {
	Points    []Point        `json:",omitempty"`
	Tracks    []Track        `json:",omitempty"`
	Tags      []Tag          `json:",omitempty"`
	files     []AttachedFile `json:",omitempty"`
	chapters  []Entry        `json:",omitempty"`
	SeekHeads []SeekHead     `json:",omitempty"`
	Info      *Info          `json:",omitempty"`
	clusters  []Cluster      `json:",omitempty"`
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
