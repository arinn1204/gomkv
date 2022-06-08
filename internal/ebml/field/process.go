package field

import (
	"encoding/binary"
	"fmt"
	"math"

	"github.com/arinn1204/gomkv/internal/array"
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
)

var read func(ebml *ebml.Ebml, data []byte) (int, error)

func init() {
	read = func(ebml *ebml.Ebml, data []byte) (int, error) {
		return ebml.File.Read(ebml.CurrPos, data)
	}
}

type FieldDataProcessor interface {
	GetFieldData(id uint32, elemSize int64, ebml *ebml.Ebml) (any, error)
}

type Processor struct {
	ebml *ebml.Ebml
}

func (p *Processor) GetFieldData(id uint32, elemSize int64) (any, error) {

	buf := make([]byte, elemSize)
	n, err := read(p.ebml, buf)

	if err != nil {
		return nil, err
	}

	p.ebml.CurrPos += int64(n)
	element := p.ebml.Specification.Data[id]

	return getElementData(buf, element)
}

func getElementData(buf []byte, element *specification.EbmlData) (any, error) {
	var err error
	var data any
	switch element.Type {
	case "binary":
		data = buf
	case "date":
		data = getData(buf)
	case "uinteger":
		data = uint(getData(buf))
	case "float":
		val := getData(buf)
		data = math.Float32frombits(uint32(val))
	case "utf-8":
		fallthrough
	case "string":
		data = string(buf)
	default:
		err = fmt.Errorf("failed to get data for %v", element.Type)
	}

	return data, err
}

func getData(buf []byte) uint64 {
	paddedBuf := make([]byte, 8)
	array.Pad(buf, paddedBuf)
	data := binary.BigEndian.Uint64(paddedBuf)
	return data
}
