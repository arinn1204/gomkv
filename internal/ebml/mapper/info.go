package mapper

import (
	"fmt"
	"io"
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
			return processInfo(id, &ebml, element, info)
		},
	)

	return info, err
}

func processInfo(id uint32, ebml *ebml.Ebml, element *specification.EbmlData, info *types.Info) error {
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
		if procErr := process(info, id, ebml); procErr != nil {
			err = utils.ConcatErr(err, procErr)
		}
	case "SegmentFamily":
		fallthrough
	case "SegmentUID":
		fallthrough
	case "NextUID":
		fallthrough
	case "PrevUID":
		elementSize, elemErr := getSize(ebml)
		if elemErr != nil {
			ebml.CurrPos += elementSize
			if elemErr == io.EOF {
				err = elemErr
				break
			}
			err = utils.ConcatErr(err, fmt.Errorf("failed to get size of %x", id))
		}
		buf := make([]byte, elementSize)
		n, _ := read(ebml, buf)
		ebml.CurrPos += int64(n)
		val, uuidErr := uuid.FromBytes(buf)

		if uuidErr != nil {
			err = utils.ConcatErr(err, uuidErr)
			break
		}

		reflect.ValueOf(info).Elem().FieldByName(element.Name).Set(reflect.ValueOf(&val))
	default:
		elementSize, sizeErr := getSize(ebml)
		if sizeErr != nil {
			err = utils.ConcatErr(err, fmt.Errorf("failed to get size of %x", id))
			if sizeErr == io.EOF {
				break
			}
		}

		ebml.CurrPos += elementSize
	}

	return err
}
