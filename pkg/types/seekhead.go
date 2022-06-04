package types

//SeekHead is the container for all the seek information
type SeekHead struct {
	Seeks []Seek
}

//Seek contains the location of other elements in the ebml document
type Seek struct {
	SeekPosition uint
	SeekID       uint64
}
