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
