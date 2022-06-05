package types

//SeekHead is the container for all the seek information
type SeekHead struct {
	Seeks []*Seek `json:",omitempty"`
}

//Seek contains the location of other elements in the ebml document
type Seek struct {
	SeekPosition uint   `json:",omitempty"`
	SeekID       uint64 `json:",omitempty"`
}
