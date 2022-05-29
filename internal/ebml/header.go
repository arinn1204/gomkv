package ebml

import (
	"encoding/binary"

	"github.com/arinn1204/gomkv/pkg/types"
)

func createHeader(ebml Ebml) (types.Header, error) {
	startPos := ebml.CurrPos
	headerSize := ebml.GetSize()

	header := types.Header{}

	for ebml.CurrPos < startPos+headerSize {
		id, err := getId(&ebml)

		if err != nil {
			return header, err
		}

		process(&header, id, &ebml)
	}

	return header, nil
}

func getId(ebml *Ebml) (uint16, error) {
	buf := make([]byte, 2)
	n, err := ebml.File.Read(ebml.CurrPos, buf)
	if err != nil {
		return 0, err
	}

	ebml.CurrPos += int64(n)

	return binary.BigEndian.Uint16(buf), nil
}

func process(header *types.Header, id uint16, ebml *Ebml) {
	// elemSize := ebml.GetSize()
	// elem := spec.Data[uint32(id)]
}
