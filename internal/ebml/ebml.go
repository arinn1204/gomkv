package ebml

import (
	"github.com/arinn1204/gomkv/internal/ebml/types"
	"github.com/arinn1204/gomkv/internal/io"
)

//EbmlReader is the interface that will parse the file to obtain the metadata
type EbmlReader interface {
	read(*io.File) types.EbmlDocument
}

//Ebml is the struct used to receive the interface
type Ebml struct{}

var ebml Ebml

func init() {
	ebml = Ebml{}
}

func (ebml Ebml) Read(f *io.File) types.EbmlDocument {
	return types.EbmlDocument{}
}
