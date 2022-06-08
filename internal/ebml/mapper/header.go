package mapper

import (
	"encoding/binary"
	"fmt"
	"io"

	"github.com/arinn1204/gomkv/internal/array"
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

type Header struct{}

func (Header) Map(size int64, ebml ebml.Ebml) (*types.Header, error) {
	header := new(types.Header)

	err := readUntil(
		&ebml,
		ebml.CurrPos+size,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			var set func(*types.Header, any)
			var err error

			switch element.Name {
			case "DocType":
				set = func(h *types.Header, a any) {
					h.DocType = a.(string)
				}
			case "DocTypeReadVersion":
				set = func(h *types.Header, a any) {
					h.DocTypeReadVersion = a.(uint)
				}
			case "DocTypeVersion":
				set = func(h *types.Header, a any) {
					h.DocTypeVersion = a.(uint)
				}
			case "EBMLMaxIDLength":
				set = func(h *types.Header, a any) {
					h.EBMLMaxIDLength = a.(uint)
				}
			case "EBMLMaxSizeLength":
				set = func(h *types.Header, a any) {
					h.EBMLMaxSizeLength = a.(uint)
				}
			case "EBMLReadVersion":
				set = func(h *types.Header, a any) {
					h.EBMLReadVersion = a.(uint)
				}
			case "EBMLVersion":
				set = func(h *types.Header, a any) {
					h.EBMLVersion = a.(uint)
				}
			default:
				ebml.CurrPos = endPos
			}
			if set != nil {
				var data any
				data, err = getFieldData(id, endPos-ebml.CurrPos, &ebml)
				set(header, data)
			}
			return err
		},
	)

	return header, err
}

//GetID is a function that will return the ID of the following EBML element
func getID(ebml *ebml.Ebml, maxCount int) (uint32, error) {
	buf := make([]byte, maxCount)
	byteToRead := 1

	var id uint32

	for byteToRead <= maxCount {
		_, err := read(ebml, buf[maxCount-byteToRead:maxCount])
		if err != nil {
			if err == io.EOF {
				return 0, err
			}
			return 0, fmt.Errorf("getID failed to read: %v", err.Error())
		}

		paddedBuf := make([]byte, 4)
		array.Pad(buf, paddedBuf)
		id = binary.BigEndian.Uint32(paddedBuf)

		if ebml.Specification.Data[id] != nil {
			break
		}

		byteToRead++
	}

	ebml.CurrPos += int64(byteToRead)

	return id, nil
}
