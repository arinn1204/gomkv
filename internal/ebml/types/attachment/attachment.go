package attachment

type AttachedFile struct {
	desc string
	name string
	mime string
	uid  uint
	data []byte
}
