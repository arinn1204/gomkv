package types

type AttachedFile struct {
	desc string
	name string
	mime string
	uid  uint
	data []byte
}

//ebml defines a single attachement with many files
type Attachment struct {
	files []AttachedFile
}
