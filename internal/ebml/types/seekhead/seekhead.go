package seekhead

import "github.com/google/uuid"

//SeekHead is the container for all the seek information
type SeekHead struct {
	seeks []Seek
}

//Seek contains the location of other elements in the ebml document
type Seek struct {
	position  uint
	elementID uuid.UUID
}
