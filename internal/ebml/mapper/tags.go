package mapper

import (
	"sync"

	"github.com/arinn1204/gomkv/internal/ebml"
	"github.com/arinn1204/gomkv/internal/ebml/specification"
	"github.com/arinn1204/gomkv/internal/utils"
	"github.com/arinn1204/gomkv/pkg/types"
)

type tags struct{}

func (tags) Map(size int64, ebmlContainer ebml.Ebml) ([]*types.Tag, error) {
	type tagContainer struct {
		tag *types.Tag
		err error
	}

	tagChan := make(chan *tagContainer, 5)
	chanCount := 0
	tags := make([]*types.Tag, 0)
	wg := &sync.WaitGroup{}

	readUntil(
		&ebmlContainer,
		ebmlContainer.CurrPos+size,
		func(id uint32, endPosition int64, element *specification.EbmlData) error {
			switch element.Name {
			case "Tag":
				wg.Add(1)
				func(e ebml.Ebml) {
					defer wg.Done()
					tag, err := processTag(&e, endPosition)

					tagChan <- &tagContainer{
						tag: tag,
						err: err,
					}
					chanCount++

				}(ebmlContainer)
			default:
				ebmlContainer.CurrPos = endPosition
			}

			return nil
		},
	)

	wg.Wait()
	var err error
	for chanCount > 0 {
		container := <-tagChan
		tags = append(tags, container.tag)
		err = utils.ConcatErr(err, container.err)
		chanCount--
	}

	return tags, err
}

func processTag(ebml *ebml.Ebml, endPosition int64) (*types.Tag, error) {
	tag := new(types.Tag)
	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPosition int64, element *specification.EbmlData) error {
			var err error
			switch element.Name {
			case "Targets":
				tag.Targets, err = processTargets(ebml, endPosition)
			case "SimpleTag":
				var simpleTag *types.SimpleTag
				simpleTag, err = processSimpleTag(ebml, endPosition)
				tag.Tags = append(tag.Tags, simpleTag)
			default:
				ebml.CurrPos = endPosition
			}

			return err
		},
	)

	return tag, err
}

func processTargets(ebml *ebml.Ebml, endPosition int64) (*types.Target, error) {
	target := new(types.Target)
	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPosition int64, element *specification.EbmlData) error {
			var err error
			switch element.Name {
			case "TargetTypeValue":
				fallthrough
			case "TargetType":
				fallthrough
			case "TagTrackUID":
				fallthrough
			case "TagEditionUID":
				fallthrough
			case "TagAttachmentUID":
				err = process(target, id, endPosition-ebml.CurrPos, ebml)
			default:
				ebml.CurrPos = endPosition
			}
			return err
		})
	return target, err
}

func processSimpleTag(ebml *ebml.Ebml, endPosition int64) (*types.SimpleTag, error) {
	tag := new(types.SimpleTag)

	wg := &sync.WaitGroup{}
	errChan := make(chan error)
	errCount := 0

	err := readUntil(
		ebml,
		endPosition,
		func(id uint32, endPosition int64, element *specification.EbmlData) error {
			switch element.Name {
			case "TagName":
				fallthrough
			case "TagLanguage":
				fallthrough
			case "TagLanguageIETF":
				fallthrough
			case "TagString":
				fallthrough
			case "TagDefault":
				fallthrough
			case "TagDefaultBogus":
				return process(tag, id, endPosition-ebml.CurrPos, ebml)
			case "SimpleTag":
				wg.Add(1)
				go func() {
					defer wg.Done()
					simpleTag, err := processSimpleTag(ebml.Copy(), endPosition)
					tag.Child = simpleTag
					errChan <- err
					errCount++
				}()
			default:
				ebml.CurrPos = endPosition
			}
			return nil
		},
	)

	for errCount > 0 {
		err = utils.ConcatErr(err, <-errChan)
	}

	return tag, err
}
