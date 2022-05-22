package seekhead

import "github.com/google/uuid"

type SeekHead struct {
	seeks []Seek
}

type Seek struct {
	position  uint
	elementId uuid.UUID
}
