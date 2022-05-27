package ebml

import (
	"github.com/arinn1204/gomkv/internal/io"
)

//EbmlReader will contain the IoReader as well as the current position of this members stream
type Ebml struct {
	File    io.Reader
	CurrPos uint
}
