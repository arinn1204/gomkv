package mapper

import (
	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
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
	entry := new(types.Entry)

	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPos int64, element *specification.EbmlData) error {
			ebml.CurrPos = endPos
			return nil
		},
	)

	return entry, err
}
