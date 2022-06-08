package mapper

import (
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
	var set func(*types.Info, any)
	switch element.Name {
	case "Duration":
		set = func(v *types.Info, a any) {
			v.Duration = a.(float32)
		}
	case "MuxingApp":
		set = func(v *types.Info, a any) {
			v.MuxingApp = a.(string)
		}
	case "WritingApp":
		set = func(v *types.Info, a any) {
			v.WritingApp = a.(string)
		}
	case "TimestampScale":
		set = func(v *types.Info, a any) {
			v.TimestampScale = a.(uint)
		}
	case "SegmentFamily":
		set = func(v *types.Info, a any) {
			var val uuid.UUID
			val, err = uuid.FromBytes(a.([]byte))
			v.SegmentFamily = &val
		}
	case "SegmentUID":
		set = func(v *types.Info, a any) {
			var val uuid.UUID
			val, err = uuid.FromBytes(a.([]byte))
			v.SegmentUID = &val
		}
	case "NextUID":
		set = func(v *types.Info, a any) {
			var val uuid.UUID
			val, err = uuid.FromBytes(a.([]byte))
			v.NextUID = &val
		}
	case "PrevUID":
		set = func(v *types.Info, a any) {
			var val uuid.UUID
			val, err = uuid.FromBytes(a.([]byte))
			v.PrevUID = &val
		}
	case "DateUTC":
		set = func(v *types.Info, a any) {
			v.DateUTC = a.(uint64)
		}

	default:
		ebml.CurrPos += size
	}

	if set != nil {
		data, fieldErr := getFieldData(id, size, ebml)
		err = utils.ConcatErr(err, fieldErr)
		set(info, data)
	}

	return err
}
