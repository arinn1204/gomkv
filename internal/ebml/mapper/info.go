package mapper

import (
	"fmt"
	"io"
	"reflect"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/utils"
	"github.com/arinn1204/gomkv/pkg/types"
	"github.com/google/uuid"
)

type info struct{}

func (info) Map(size int64, ebml ebml.Ebml) (*types.Info, error) {
	info := types.Info{}
	infoEnd := ebml.CurrPos + size
	var err error

	for ebml.CurrPos < infoEnd {
		if err == io.EOF {
			break
		}
		id, elemErr := GetID(&ebml, 3)

		if elemErr != nil {
			err = utils.ConcatErr(err, elemErr)
			break
		}

		element := ebml.Specification.Data[id]
		if element == nil {
			err = utils.ConcatErr(err, fmt.Errorf("unrecognized id of 0x%X", id))
			break
		}

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
			if procErr := process(&info, id, &ebml); procErr != nil {
				err = utils.ConcatErr(err, procErr)
				if procErr == io.EOF {
					break
				}
			}
		case "SegmentFamily":
			fallthrough
		case "SegmentUID":
			fallthrough
		case "NextUID":
			fallthrough
		case "PrevUID":
			elementSize, elemErr := getSize(&ebml)
			if elemErr != nil {
				ebml.CurrPos += elementSize
				if elemErr == io.EOF {
					err = elemErr
					break
				}
				err = utils.ConcatErr(err, fmt.Errorf("failed to get size of %x", id))
				continue
			}
			buf := make([]byte, elementSize)
			n, _ := read(&ebml, buf)
			ebml.CurrPos += int64(n)
			val, uuidErr := uuid.FromBytes(buf)

			if uuidErr != nil {
				err = utils.ConcatErr(err, uuidErr)
				continue
			}

			reflect.ValueOf(&info).Elem().FieldByName(element.Name).Set(reflect.ValueOf(val))
		default:
			elementSize, sizeErr := getSize(&ebml)
			if sizeErr != nil {
				err = utils.ConcatErr(err, fmt.Errorf("failed to get size of %x", id))
				if sizeErr == io.EOF {
					break
				}
				continue
			}

			ebml.CurrPos += elementSize
		}

	}

	return &info, err
}
