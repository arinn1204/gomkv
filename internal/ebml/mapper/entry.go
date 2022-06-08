package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

func makeTrackEntry(ebml *ebml.Ebml, endPosition int64) (*types.Entry, error) {
	entry := new(types.Entry)

	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			var err error
			var set func(*types.Entry, any)
			switch element.Name {
			case "Name":
				if set == nil {
					set = func(v *types.Entry, a any) {
						v.Name = a.(string)
					}
				}
				fallthrough
			case "TrackUID":
				if set == nil {
					set = func(v *types.Entry, a any) {
						v.TrackUID = a.(uint)
					}
				}
				fallthrough
			case "Language":
				if set == nil {
					set = func(v *types.Entry, a any) {
						v.Language = a.(string)
					}
				}
				fallthrough
			case "LanguageIETF":
				if set == nil {
					set = func(v *types.Entry, a any) {
						v.LanguageIETF = a.(string)
					}
				}
				fallthrough
			case "CodecID":
				if set == nil {
					set = func(v *types.Entry, a any) {
						v.CodecID = a.(string)
					}
				}
				fallthrough
			case "CodecName":
				if set == nil {
					set = func(v *types.Entry, a any) {
						v.CodecName = a.(string)
					}
				}
				var data any
				data, err = getFieldData(id, endPos-ebml.CurrPos, ebml)
				set(entry, data)
			case "Video":
				var video *types.Video
				video, err = makeVideoEntry(ebml, endPos)
				entry.Video = video
			case "Audio":
				var audio *types.Audio
				audio, err = makeAudioEntry(ebml, endPos)
				entry.Audio = audio
			default:
				ebml.CurrPos = endPos
			}
			return err
		},
	)

	return entry, err
}
