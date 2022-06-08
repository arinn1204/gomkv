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
				go func(e ebml.Ebml) {
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
			var set func(*types.Target, any)
			var err error
			switch element.Name {
			case "TargetTypeValue":
				set = func(v *types.Target, a any) {
					v.TargetTypeValue = a.(uint)
				}
			case "TargetType":
				set = func(v *types.Target, a any) {
					v.TargetType = a.(string)
				}
			case "TagTrackUID":
				set = func(v *types.Target, a any) {
					v.TagTrackUID = a.(uint)
				}
			case "TagEditionUID":
				set = func(v *types.Target, a any) {
					v.TagEditionUID = a.(uint)
				}
			case "TagAttachmentUID":
				set = func(v *types.Target, a any) {
					v.TagAttachmentUID = a.(uint)
				}
			default:
				ebml.CurrPos = endPosition
			}
			if set != nil {
				var data any
				data, err = getFieldData(id, endPosition-ebml.CurrPos, ebml)
				set(target, data)
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
			var set func(*types.SimpleTag, any)
			var err error
			switch element.Name {
			case "TagName":
				set = func(v *types.SimpleTag, a any) {
					v.TagName = a.(string)
				}
			case "TagLanguage":
				set = func(v *types.SimpleTag, a any) {
					v.TagLanguage = a.(string)
				}
			case "TagLanguageIETF":
				set = func(v *types.SimpleTag, a any) {
					v.TagLanguageIETF = a.(string)
				}
			case "TagString":
				set = func(v *types.SimpleTag, a any) {
					v.TagString = a.(string)
				}
			case "TagDefault":
				set = func(v *types.SimpleTag, a any) {
					v.TagDefault = a.(uint)
				}
			case "TagDefaultBogus":
				set = func(v *types.SimpleTag, a any) {
					v.TagDefaultBogus = a.(string)
				}
			case "SimpleTag":
				wg.Add(1)
				go func() {
					defer wg.Done()
					simpleTag, err := processSimpleTag(ebml.Copy(), endPosition)
					tag.Child = simpleTag
					errChan <- err
					errCount++
				}()
				fallthrough
			default:
				ebml.CurrPos = endPosition
			}
			if set != nil {
				var data any
				data, err = getFieldData(id, endPosition-ebml.CurrPos, ebml)
				set(tag, data)
			}
			return err
		},
	)

	for errCount > 0 {
		err = utils.ConcatErr(err, <-errChan)
	}

	return tag, err
}
