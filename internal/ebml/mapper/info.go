package mapper

import (
	"fmt"
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
		id, elemErr := GetID(&ebml, 3)

		if elemErr != nil {
			err = utils.ConcatErr(err, elemErr)
			continue
		}

		element := ebml.Specification.Data[id]
		if element == nil {
			err = utils.ConcatErr(err, fmt.Errorf("unrecognized id of %x", id))
			continue
		}

		switch element.Name {
		case "SegmentFamily":
			fallthrough
		case "SegmentUID":
			fallthrough
		case "NextUID":
			fallthrough
		case "PrevUID":
			elementSize, elemErr := getSize(&ebml)
			if elemErr != nil {
				err = utils.ConcatErr(err, fmt.Errorf("failed to get size of %x", id))
				continue
			}
			buf := make([]byte, elementSize)
			n, _ := read(&ebml, buf)
			ebml.CurrPos += int64(n)
			val, uuidErr := uuid.FromBytes(buf)

			if uuidErr != nil {
				err = utils.ConcatErr(err, uuidErr)
			}

			reflect.ValueOf(&info).Elem().FieldByName(element.Name).Set(reflect.ValueOf(val))
		case "Translations":
			//todo fill me out
		default:
			process(&info, id, &ebml)
		}

	}

	return &info, err
}
