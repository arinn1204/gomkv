package ebml

import (
	"github.com/arinn1204/gomkv/internal/ebml/types"
	"github.com/arinn1204/gomkv/internal/filesystem"
)

//EbmlReader will contain the IoReader as well as the current position of this members stream
type Ebml struct {
	File    filesystem.Reader
	CurrPos uint
}

func (ebml Ebml) Read() types.EbmlDocument {
	return types.EbmlDocument{}
}
