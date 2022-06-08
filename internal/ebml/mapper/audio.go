package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/pkg/types"
)

func makeAudioEntry(ebml *ebml.Ebml, endPosition int64) (*types.Audio, error) {
	audio := new(types.Audio)
	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPosition int64, element *specification.EbmlData) error {
			var set func(*types.Audio, any)
			var err error
			switch element.Name {
			case "SamplingFrequency":
				set = func(a1 *types.Audio, a2 any) {
					a1.SamplingFrequency = a2.(float32)
				}
				fallthrough
			case "OutputSamplingFrequency":
				if set == nil {
					set = func(a1 *types.Audio, a2 any) {
						a1.OutputSamplingFrequency = a2.(float32)
					}
				}
				fallthrough
			case "Channels":
				if set == nil {
					set = func(a1 *types.Audio, a2 any) {
						a1.Channels = a2.(uint)
					}
				}
				fallthrough
			case "BitDepth":
				if set == nil {
					set = func(a1 *types.Audio, a2 any) {
						a1.BitDepth = a2.(uint)
					}
				}
				var data any
				data, err = getFieldData(id, endPosition-ebml.CurrPos, ebml)
				set(audio, data)
			default:
				ebml.CurrPos = endPosition
			}
			return err
		},
	)
	return audio, err
}
