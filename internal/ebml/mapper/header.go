package mapper

import (
	"encoding/binary"

	"github.com/arinn1204/gomkv/internal/array"
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/pkg/ebml/types"
)

func Map(ebml ebml.Ebml) (types.Header, error) {
	startPos := ebml.CurrPos
	headerSize := ebml.GetSize()

	header := types.Header{}

	for ebml.CurrPos < startPos+headerSize {
		elementSize := ebml.GetSize()
		buf := make([]byte, elementSize)
		paddedBuf := make([]byte, 8)
		err := array.Pad(buf, paddedBuf)
		if err != nil {
			return types.Header{}, err
		}

		id := binary.BigEndian.Uint64(paddedBuf)
		process(&header, id, &ebml)
	}

	return header, nil
}

func process(header *types.Header, id uint64, ebml *ebml.Ebml) {

}
