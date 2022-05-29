package ebml

import (
	"github.com/arinn1204/gomkv/internal/ebml/types"
	"github.com/arinn1204/gomkv/internal/filesystem"
)

//Ebml will contain the IoReader as well as the current position of this members stream
type Ebml struct {
	File    filesystem.Reader
	CurrPos int64
}

func (ebml Ebml) Read() types.EbmlDocument {
	return types.EbmlDocument{
		Header: getHeader(ebml),
	}
}

func getHeader(ebml Ebml) types.Header {
	header := types.Header{}
	size := ebml.GetSize()

	endPos := ebml.CurrPos + size
	_ = endPos

	return header
}
