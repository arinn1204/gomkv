package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/utils"
	"github.com/arinn1204/gomkv/pkg/types"
)

type tracks struct{}

func (tracks) Map(size int64, ebml ebml.Ebml) (*types.DisplayTrack, error) {
	endPosition := ebml.CurrPos + size

	entries, err := mapTracks(&ebml, endPosition)

	return &types.DisplayTrack{Entries: entries}, err
}

func mapTracks(ebml *ebml.Ebml, endPosition int64) ([]*types.Entry, error) {
	entries := make([]*types.Entry, 0)

	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			entry, err := makeTrackEntry(ebml, endPos)
			entries = append(entries, entry)
			return err
		},
	)

	return entries, err
}

func makeTrackEntry(ebml *ebml.Ebml, endPosition int64) (*types.Entry, error) {
	errChan := make(chan error, 7)
	entry := new(types.Entry)
	entries := 0
	readUntil(
		ebml,
		endPosition,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			switch element.Name {
			case "Name":
				fallthrough
			case "Language":
				fallthrough
			case "LanguageIETF":
				fallthrough
			case "CodecID":
				fallthrough
			case "CodecName":
				go func() {
					entries++
					errChan <- process(entry, id, endPos-ebml.CurrPos, ebml)
				}()
			case "Video":
				go func() {
					entries++
					video, err := makeVideoEntry(*ebml, endPos)
					entry.Video = video

					errChan <- err
				}()
			case "Audio":
				go func() {
					entries++
					audio, err := makeAudioEntry(*ebml, endPos)
					entry.Audio = audio

					errChan <- err
				}()
			default:
				ebml.CurrPos = endPos
			}
			return nil
		},
	)

	var err error
	for i := 0; i < entries; i++ {
		err = utils.ConcatErr(err, <-errChan)
	}

	close(errChan)
	return entry, err
}

func makeVideoEntry(ebml ebml.Ebml, endPosition int64) (*types.Video, error) {
	return nil, nil
}

func makeAudioEntry(ebml ebml.Ebml, endPosition int64) (*types.Audio, error) {
	return nil, nil
}
