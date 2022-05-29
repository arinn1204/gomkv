package ebml

import (
	"encoding/binary"
	"fmt"

	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/filesystem"
	"github.com/arinn1204/gomkv/pkg/types"
)

//Ebml will contain the IoReader as well as the current position of this members stream
type Ebml struct {
	File              filesystem.Reader
	CurrPos           int64
	SpecificationPath string
}

func (ebml Ebml) Read() (types.EbmlDocument, error) {
	spec := specification.GetSpecification(ebml.SpecificationPath)
	err := validateMagicNum(&ebml, spec)
	if err != nil {
		return types.EbmlDocument{}, err
	}

	header, _ := createHeader(ebml, spec)

	return types.EbmlDocument{
			Header:   header,
			Segments: make([]types.Segment, 1),
		},
		nil
}

func validateMagicNum(ebml *Ebml, spec specification.Ebml) error {
	idBuf := make([]byte, 4)
	n, err := ebml.File.Read(ebml.CurrPos, idBuf)

	if err != nil {
		return err
	}

	ebml.CurrPos += int64(n)

	id := binary.BigEndian.Uint32(idBuf)
	elem := spec.Data[id]

	if elem.Name != "EBML" {
		return fmt.Errorf("incorrect type of file expected magic number found %x", id)
	}

	return nil
}
