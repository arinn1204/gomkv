package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

func makeVideoEntry(ebml *ebml.Ebml, endPosition int64) (*types.Video, error) {
	video := new(types.Video)

	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPosition int64, element *specification.EbmlData) error {
			var set func(*types.Video, any)
			var err error
			switch element.Name {
			case "PixelWidth":
				if set == nil {
					set = func(v *types.Video, a any) {
						v.PixelWidth = a.(uint)
					}
				}
				fallthrough
			case "PixelHeight":
				if set == nil {
					set = func(v *types.Video, a any) {
						v.PixelHeight = a.(uint)
					}
				}
				fallthrough
			case "DisplayWidth":
				if set == nil {
					set = func(v *types.Video, a any) {
						v.DisplayWidth = a.(uint)
					}
				}
				fallthrough
			case "DisplayHeight":
				if set == nil {
					set = func(v *types.Video, a any) {
						v.DisplayHeight = a.(uint)
					}
				}
				fallthrough
			case "DisplayUnit":
				if set == nil {
					set = func(v *types.Video, a any) {
						v.DisplayUnit = a.(uint)
					}
				}
				fallthrough
			case "AspectRatioType":
				if set == nil {
					set = func(v *types.Video, a any) {
						v.AspectRatioType = a.(uint)
					}
				}
				var data any
				data, err = getFieldData(id, endPosition-ebml.CurrPos, ebml)
				set(video, data)
			default:
				ebml.CurrPos = endPosition
			}

			return err
		},
	)
	return video, err
}
