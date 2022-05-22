package tag

//Tag is a metadata descriptor
type Tag struct {
	target Target
	tags   []SimpleTag
}

//Target describes what element the tag is describing
//If target is empty then it applies to the whole segment
type Target struct {
	logicalLevelValue uint
	logicalLevel      uint
	trackUids         []uint
	editionUids       []uint
	chapterUids       []uint
	attachementUids   []uint
}

//SimpleTag contains general information about the target
type SimpleTag struct {
	child          *SimpleTag
	name           string
	language       string
	languageIETF   string
	defaultLaunage uint
	tagString      string
	binary         []byte
}
