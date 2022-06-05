package mapper

import (
	"reflect"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/utils"
	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/google/uuid"
)

type info struct{}

func (info) Map(size int64, ebml ebml.Ebml) (*types.Info, error) {
	info := new(types.Info)
	infoEnd := ebml.CurrPos + size

	err := readUntil(
		&ebml,
		infoEnd,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			return processInfo(id, endPos-ebml.CurrPos, &ebml, element, info)
		},
	)

	return info, err
}

func processInfo(id uint32, size int64, ebml *ebml.Ebml, element *specification.EbmlData, info *types.Info) error {
	var err error
	switch element.Name {
	case "Duration":
		fallthrough
	case "MuxingApp":
		fallthrough
	case "WritingApp":
		fallthrough
	case "TimestampScale":
		fallthrough
	case "DateUTC":
		if procErr := process(info, id, size, ebml); procErr != nil {
			err = utils.ConcatErr(err, procErr)
		}
	case "SegmentFamily":
		fallthrough
	case "SegmentUID":
		fallthrough
	case "NextUID":
		fallthrough
	case "PrevUID":
		err = utils.ConcatErr(err, processUUID(ebml, id, size, info, element))
	default:
		ebml.CurrPos += size
	}

	return err
}

func processUUID(ebml *ebml.Ebml, id uint32, elementSize int64, info *types.Info, element *specification.EbmlData) error {
	var err error
	buf := make([]byte, elementSize)
	n, _ := read(ebml, buf)
	ebml.CurrPos += int64(n)
	val, uuidErr := uuid.FromBytes(buf)

	if uuidErr != nil {
		return utils.ConcatErr(err, uuidErr)
	}

	reflect.ValueOf(info).Elem().FieldByName(element.Name).Set(reflect.ValueOf(&val))
	return err
}
