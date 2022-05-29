package types

//AttachedFile is one of many that may be attached to the EBML document
//ID is 0x61A7
type AttachedFile struct {
	desc string
	name string
	mime string
	uid  uint
	data []byte
	id   int `default:"0x61A7"`
}
