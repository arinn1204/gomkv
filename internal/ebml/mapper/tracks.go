package mapper

import (
	"sync"

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

func mapTracks(ebmlContainer *ebml.Ebml, endPosition int64) ([]*types.Entry, error) {
	entries := make([]*types.Entry, 0)
	type entryContainer struct {
		entry *types.Entry
		err   error
	}

	//at most 5 tracks in parallel
	entryChan := make(chan *entryContainer, 5)
	count := 0
	wg := &sync.WaitGroup{}

	err := readUntil(
		ebmlContainer,
		endPosition,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			count++
			wg.Add(1)
			go func(endPos int64, ebml ebml.Ebml) {
				defer wg.Done()

				entry, err := makeTrackEntry(&ebml, endPos)
				entryChan <- &entryContainer{
					entry: entry,
					err:   err,
				}
			}(endPos, *ebmlContainer)
			ebmlContainer.CurrPos = endPos
			return nil
		},
	)

	wg.Wait()
	for i := 0; i < count; i++ {
		container := <-entryChan
		err = utils.ConcatErr(err, container.err)
		entries = append(entries, container.entry)
	}

	return entries, err
}

func makeTrackEntry(ebml *ebml.Ebml, endPosition int64) (*types.Entry, error) {
	entry := new(types.Entry)

	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			var err error
			switch element.Name {
			case "Name":
				fallthrough
			case "TrackUID":
				fallthrough
			case "Language":
				fallthrough
			case "LanguageIETF":
				fallthrough
			case "CodecID":
				fallthrough
			case "CodecName":
				err = process(entry, id, endPos-ebml.CurrPos, ebml)
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

func makeVideoEntry(ebml *ebml.Ebml, endPosition int64) (*types.Video, error) {
	video := new(types.Video)

	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPosition int64, element *specification.EbmlData) error {
			switch element.Name {
			case "PixelWidth":
				fallthrough
			case "PixelHeight":
				fallthrough
			case "DisplayWidth":
				fallthrough
			case "DisplayHeight":
				fallthrough
			case "DisplayUnit":
				fallthrough
			case "AspectRatioType":
				return process(video, id, endPosition-ebml.CurrPos, ebml)
			default:
				ebml.CurrPos = endPosition
			}

			return nil
		},
	)
	return video, err
}

func makeAudioEntry(ebml *ebml.Ebml, endPosition int64) (*types.Audio, error) {
	audio := new(types.Audio)
	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPosition int64, element *specification.EbmlData) error {
			switch element.Name {
			case "SamplingFrequency":
				fallthrough
			case "OutputSamplingFrequency":
				fallthrough
			case "Channels":
				fallthrough
			case "BitDepth":
				return process(audio, id, endPosition-ebml.CurrPos, ebml)
			default:
				ebml.CurrPos = endPosition
			}
			return nil
		},
	)
	return audio, err
}
